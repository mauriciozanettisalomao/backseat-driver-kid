package ports

import (
	"context"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
)

// Interactable defines the behavior of interactable actions
type Interactable interface {
	ExpandKnowledge(context.Context, *domain.Interaction) error
	Interact(context.Context, *domain.Interaction) error
}
