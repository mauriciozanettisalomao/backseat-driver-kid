package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/ports"

	"golang.org/x/sync/errgroup"
)

// Interaction implements the run behavior for the interact service
type Interaction struct {
	text             string
	numRoutines      int
	syncInteractions *domain.Interaction
	inputReader      ports.InputReader
	interaction      ports.Interactable
	outputWriter     ports.OutputWriter
}

func (i *Interaction) Run(ctx context.Context) error {

	interactions, err := i.inputReader.Read(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Error reading input", "error", err)
		return err
	}

	i.syncInteractions = &domain.Interaction{
		Preamble: interactions.Preamble,
		Prompts:  []*domain.Prompt{},
	}

	slog.InfoContext(ctx, "Interactions received",
		"interaction_prompts", len(interactions.Prompts),
	)

	errs, ctx := errgroup.WithContext(ctx)

	interactionsChan := make(chan *domain.Prompt, len(interactions.Prompts))

	// producer
	go func() {
		defer close(interactionsChan)
		for _, prompt := range interactions.Prompts {
			i.syncInteractions.Prompts = append(i.syncInteractions.Prompts, prompt)
			interactionsChan <- prompt
		}
	}()

	for idx := 0; idx < i.numRoutines; idx++ {

		// consumer
		errs.Go(func() (err error) {

			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("recovered for %v", r)
					slog.Error("error processing event registrations",
						"err", err,
					)
				}
			}()

			for ic := range interactionsChan {

				errInteract := i.interaction.Interact(ctx, &domain.Interaction{
					Preamble: domain.Preamble{
						Context:      i.syncInteractions.Preamble.Context,
						Instructions: i.syncInteractions.Preamble.Instructions,
						Examples:     i.syncInteractions.Preamble.Examples,
					},
					Prompts: []*domain.Prompt{ic},
				})
				if errInteract != nil {
					slog.ErrorContext(ctx, "error processing the interaction",
						"err", errInteract,
						"input", ic.Input,
					)
					ic.Output = fmt.Sprintf("error processing the interaction: %v", errInteract)
					continue
				}

			}
			return err
		})

	}

	err = errs.Wait()
	if err != nil {
		slog.ErrorContext(ctx, "error processing the interactions",
			"err", err,
		)
	}

	errOutputWriter := i.outputWriter.Write(ctx, i.syncInteractions)
	if errOutputWriter != nil {
		slog.ErrorContext(ctx, "error writing the interactions",
			"err", errOutputWriter,
		)
		return errOutputWriter
	}

	// promptContext := `Analyze the following consent text and categorize it into one of these categories:
	// 	1. Linux Foundation News
	// 	2. Project Specific News
	// 	3. Education (training and certification)
	// 	4. Unknown`

	// instructions := `Respond only with the category, e.g.: 1. Linux Foundation News.`

	// interactionData := &domain.Interaction{
	// 	Preamble: domain.Preamble{
	// 		Context:      promptContext,
	// 		Instructions: instructions,
	// 	},
	// 	Input: []string{i.text},
	// }
	// err := i.interaction.Interact(ctx, interactionData)
	// if err != nil {
	// 	return err
	// }

	// response := strings.TrimSpace(interactionData.Output)
	// slog.InfoContext(ctx, "Response received", "response", response)

	return nil

}

// NewInteractServiceService creates a new interact service
func NewInteractServiceService() Handler {
	return &Interaction{}
}
