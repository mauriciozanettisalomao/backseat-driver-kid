// Copyright (c) 2025 Maurício Zanetti Salomão
// Licensed under the MIT License. See the LICENSE file for details.

package ports

import (
	"context"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
)

// InputReader defines the behavior of an input reader
type InputReader interface {
	Read(ctx context.Context) (*domain.Interaction, error)
}
