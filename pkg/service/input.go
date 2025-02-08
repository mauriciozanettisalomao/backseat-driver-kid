// Copyright (c) 2025 Maurício Zanetti Salomão
// Licensed under the MIT License. See the LICENSE file for details.

package service

import (
	"context"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/ports"
)

type InputService struct {
	inputReader ports.InputReader
}

func (i *InputService) Read(ctx context.Context) (interface{}, error) {
	return i.inputReader.Read(ctx)
}

func NewInputService(inputReader ports.InputReader) *InputService {
	return &InputService{
		inputReader: inputReader,
	}
}
