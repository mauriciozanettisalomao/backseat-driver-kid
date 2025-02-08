// Copyright (c) 2025 Maurício Zanetti Salomão
// Licensed under the MIT License. See the LICENSE file for details.

package file

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/misc/formatter"
)

const tmplAnalysis = `## Query Analysis

{{- if .Preamble.Context }}
### Context
{{ .Preamble.Context }}
{{ end }}
Given the additional knowledge provided [here](.ExtendedKownledgeDir), here is the analysis of the queries:

### Prompt
<ol>
{{- range .Prompts }}
  <li>{{ .Input }}</li>
	  {{ .Output }}

{{ end }}
</ol>

`

type fileOutput struct {
	file string
}

func (f *fileOutput) Write(ctx context.Context, i *domain.Interaction) error {

	analysis, err := formatter.ParseTemplate(ctx, "analysis", tmplAnalysis, i.Interaction)
	if err != nil {
		return fmt.Errorf("error parsing template: %w", err)
	}
	return f.createFile(ctx, f.file, analysis)
}

func (f *fileOutput) createFile(ctx context.Context, filename string, content string) error {

	file, errCreate := os.Create(filename)
	if errCreate != nil {
		slog.ErrorContext(ctx, "error creating file",
			"error", errCreate,
			"filename", filename,
		)
		return errCreate
	}
	defer file.Close()

	_, errWrite := file.WriteString(content)
	if errWrite != nil {
		slog.ErrorContext(ctx, "error writing to file",
			"error", errWrite,
			"filename", filename,
		)
		return errWrite
	}

	return nil
}

// NewOutputWriter creates a new file output
func NewOutputWriter(file string) *fileOutput {
	return &fileOutput{
		file: file,
	}
}
