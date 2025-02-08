// Copyright (c) 2025 Maurício Zanetti Salomão
// Licensed under the MIT License. See the LICENSE file for details.

package llm

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"time"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/misc/formatter"
	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/ports"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

const tmpl = `{{- if .Preamble.Context }}
{{ .Preamble.Context }}
{{ end }}

{{- if .ExtendedKnowledge }}
Reference Information:
{{ .ExtendedKnowledge }}
{{ end }}

{{- if .Preamble.Examples }}
Examples:
{{ .Preamble.Examples }}
{{ end }}

{{- range .Prompts }}
*** USER QUERY ***: {{ .Input }}
{{ end }}
 
{{- if .Preamble.Instructions }}
Directives: {{ .Preamble.Instructions }}
{{ end }}`

type ollamaLLM struct {
	ExtendedKnowledgeContent string
	llm                      *ollama.LLM
}

func (o *ollamaLLM) ExpandKnowledge(ctx context.Context, i *domain.Interaction) error {

	// TODO: Implement RAG using an in-memory vector store or Postgres (pgvector).
	// - **In-Memory**: Use Qdrant (memory mode) or Chroma for quick retrieval without persistence.
	// - **Postgres (pgvector)**: Store embeddings in a `VECTOR` column and use `<->` for similarity search.
	// - Consider FAISS for optimized nearest neighbor search if needed.

	// for now, we just extend the knowledge adding the document as more context
	// simulating a real expansion of the knowledge base (limited by the model)

	i.Interaction.ExtendedKnowledge = string(i.Interaction.ExtendedKnowledgeContent)

	return nil
}

func (o *ollamaLLM) Interact(ctx context.Context, i *domain.Interaction) error {

	start := time.Now()
	defer func() {
		slog.DebugContext(ctx, "interaction time in milliseconds",
			"elapsedTime", time.Since(start).Milliseconds())
	}()

	query, err := formatter.ParseTemplate(ctx, "ollamaPrompt", tmpl, i.Interaction)
	if err != nil {
		slog.ErrorContext(ctx, "error parsing prompt", "error", err)
		return err
	}

	slog.DebugContext(ctx, "query to llm started", "query", query)
	completion, err := llms.GenerateFromSinglePrompt(ctx, o.llm, query)
	if err != nil {
		slog.ErrorContext(ctx, "error calling llm", "error", err)
		return err
	}

	// we always have only one prompt for each interaction and, once all prompts are answered,
	// we can output the completion merging with the prompt
	if len(i.Interaction.Prompts) != 1 {
		slog.ErrorContext(ctx, "error processing the interaction", "error", err)
		return errors.New("invalid number of prompts")
	}
	i.Interaction.Prompts[0].Output = completion

	return nil
}

// NewLLM creates a new LLM instance
func NewLLM(model string) ports.Interactable {

	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal("no llm configured")
	}

	return &ollamaLLM{
		llm: llm,
	}
}
