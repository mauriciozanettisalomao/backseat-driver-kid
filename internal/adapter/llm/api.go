package llm

import (
	"context"
	"log"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/ports"

	"github.com/tmc/langchaingo/llms/ollama"
)

type llmAPI struct {
	llm *ollama.LLM
}

func (l *llmAPI) Interact(context.Context, *domain.Interaction) error {

	return nil
}

// NewLLMAPI creates a new LLM API
func NewLLMAPI(model string) ports.Interactable {

	llm, err := ollama.New(ollama.WithModel(model))
	if err != nil {
		log.Fatal("no llm configured")
	}

	return &llmAPI{
		llm: llm,
	}
}
