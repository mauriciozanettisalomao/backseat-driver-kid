package ports

import (
	"context"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
)

// OutputWriter defines the behavior of an output writer
type OutputWriter interface {
	Write(context.Context, *domain.Interaction) error
}
