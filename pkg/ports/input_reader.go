package ports

import (
	"context"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
)

// InputReader defines the behavior of an input reader
type InputReader interface {
	Read(ctx context.Context) (*domain.Interaction, error)
}
