package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

// promptCmd represents the prompt command
var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Prompt commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("please inform the subcommand")
	},
}

func init() {
	rootCmd.AddCommand(promptCmd)
}
