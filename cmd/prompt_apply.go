// Copyright (c) 2025 Maurício Zanetti Salomão
// Licensed under the MIT License. See the LICENSE file for details.

package cmd

import (
	"context"
	"log/slog"

	readerWriterAdapter "github.com/mauriciozanettisalomao/backseat-driver-kid/internal/adapter/file"
	interactAdapter "github.com/mauriciozanettisalomao/backseat-driver-kid/internal/adapter/llm"
	logging "github.com/mauriciozanettisalomao/backseat-driver-kid/log"
	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/service"

	"github.com/spf13/cobra"
)

const (
	defaultNumRoutines = 1
)

// cventImportAttendeesCmd represents the cventImportAttendeesCmd command
var promptApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply the prompt based on the input",
	Long: `This command will apply the prompt based on the input which should be a configuration file.

	This configuration file should be a yaml file with the following structure:

	apiVersion: v1
		kind: ConfigMap
		metadata:
		name: interaction-config
		data:
		extendedKownledgeDir:
			- "input/extended-knowledge"
		interaction:
			preamble:
			context: "This is a simulated interaction between a user and a system."
			instructions: "Provide clear and concise answers based on the context provided."
			examples: |
				Example 1: User asks about system health, and the response should indicate the status.
				Example 2: User asks for an operation status, and the system should provide a detailed report.
			prompts:
			- input: "What is the status of the system?"
			promptFile: "input/questions.txt"`,

	RunE: func(cmd *cobra.Command, args []string) error {

		logging.InitStructureLogConfig()

		ctx, cancel := context.WithCancel(cmd.Context())
		defer cancel()

		// flags
		numRoutines, errNumRoutines := cmd.Flags().GetInt("routines")
		if errNumRoutines != nil {
			slog.ErrorContext(ctx, "Error getting the number of routines", "error", errNumRoutines)
			numRoutines = defaultNumRoutines
		}

		inputConfig, errInputConfig := cmd.Flags().GetString("input")
		if errInputConfig != nil {
			slog.ErrorContext(ctx, "Error getting the input file", "error", errInputConfig)
			inputConfig = "input/interaction-config.yaml"
		}

		model, errModel := cmd.Flags().GetString("model")
		if errModel != nil {
			slog.ErrorContext(ctx, "Error getting the model", "error", errModel)
			model = "llama"
		}

		output, errOutput := cmd.Flags().GetString("output")
		if errOutput != nil {
			slog.ErrorContext(ctx, "Error getting the output file", "error", errOutput)
			output = "output/interaction-output.yaml"
		}

		service.NewInteract(
			service.WithNumRoutines(numRoutines),
			service.WithReader(readerWriterAdapter.NewInputReader(inputConfig)),
			service.WithInteractable(interactAdapter.NewLLM(model)),
			service.WithWriter(readerWriterAdapter.NewOutputWriter(output)),
		).Run(ctx)

		return nil
	},
}

func init() {
	promptCmd.AddCommand(promptApplyCmd)

	promptApplyCmd.Flags().Int("routines", 1, "Number of go routines to ask the questions concurrently")
	promptApplyCmd.Flags().String("model", "llama2", "Model to be used to interact with the user")
	promptApplyCmd.Flags().String("input", "resources/input/interaction-config.yaml", "Input file with the interaction configuration")
	promptApplyCmd.Flags().String("output", "resources/output/analysis.md", "Output file with the interaction output")
}
