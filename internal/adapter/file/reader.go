package file

import (
	"context"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/ports"
)

type fileInput struct {
	file string
}

func (f *fileInput) Read(ctx context.Context) (*domain.Interaction, error) {

	return nil, nil

}

// NewInputReader creates a new file input reader
func NewInputReader(file string) ports.InputReader {
	return &fileInput{file}
}
