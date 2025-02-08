// Copyright (c) 2025 Maurício Zanetti Salomão
// Licensed under the MIT License. See the LICENSE file for details.

package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/models"
	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/ports"

	"golang.org/x/sync/errgroup"
)

// Options helper function to ser build the orchestrator
type Options func(*interaction)

// WithNumRoutines is an option to set the number of go routines
func WithNumRoutines(numRoutines int) func(*interaction) {
	return func(a *interaction) {
		a.numRoutines = numRoutines
	}
}

// WithReader is an option to set the input reader
func WithReader(reader ports.InputReader) func(*interaction) {
	return func(a *interaction) {
		a.inputReader = reader
	}
}

// WithInteractable is an option to set the interactable
func WithInteractable(interactable ports.Interactable) func(*interaction) {
	return func(a *interaction) {
		a.interaction = interactable
	}
}

// WithWriter is an option to set the output writer
func WithWriter(writer ports.OutputWriter) func(*interaction) {
	return func(a *interaction) {
		a.outputWriter = writer
	}
}

// Interaction implements the run behavior for the interact service
type interaction struct {
	text             string
	numRoutines      int
	syncInteractions *domain.Interaction
	inputReader      ports.InputReader
	interaction      ports.Interactable
	outputWriter     ports.OutputWriter
}

func (i *interaction) Run(ctx context.Context) error {

	interactions, err := i.inputReader.Read(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Error reading input", "error", err)
		return err
	}

	i.syncInteractions = &domain.Interaction{
		Interaction: &models.Interaction{
			Preamble:                 interactions.Preamble,
			Prompts:                  []*models.Prompt{},
			ExtendedKnowledgeContent: interactions.ExtendedKnowledgeContent,
		},
	}

	slog.InfoContext(ctx, "Interactions received",
		"interaction_prompts", len(interactions.Prompts),
	)

	errs, ctx := errgroup.WithContext(ctx)

	interactionsChan := make(chan *models.Prompt, len(interactions.Prompts))

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

			errExpandKnowledge := i.interaction.ExpandKnowledge(ctx, i.syncInteractions)
			if errExpandKnowledge != nil {
				slog.ErrorContext(ctx, "error expanding the knowledge",
					"err", errExpandKnowledge,
				)
				return errExpandKnowledge
			}

			for ic := range interactionsChan {

				errInteract := i.interaction.Interact(ctx, &domain.Interaction{
					Interaction: &models.Interaction{
						Preamble: &models.Preamble{
							Context:      i.syncInteractions.Preamble.Context,
							Instructions: i.syncInteractions.Preamble.Instructions,
							Examples:     i.syncInteractions.Preamble.Examples,
						},
						Prompts:           []*models.Prompt{ic},
						ExtendedKnowledge: i.syncInteractions.ExtendedKnowledge,
					},
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

	return nil

}

// NewInteract creates a new interact service
func NewInteract(opts ...Options) Handler {
	i := &interaction{}
	for _, opt := range opts {
		opt(i)
	}
	return i
}
