// Copyright (c) 2025 Maurício Zanetti Salomão
// Licensed under the MIT License. See the LICENSE file for details.

package ports

import (
	"context"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
)

// OutputWriter defines the behavior of an output writer
type OutputWriter interface {
	Write(context.Context, *domain.Interaction) error
}
