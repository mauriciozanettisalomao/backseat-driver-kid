// Copyright (c) 2025 Maurício Zanetti Salomão
// Licensed under the MIT License. See the LICENSE file for details.

package service

import "context"

// Handler defines the behavior of the main method
type Handler interface {
	Run(ctx context.Context) error
}
