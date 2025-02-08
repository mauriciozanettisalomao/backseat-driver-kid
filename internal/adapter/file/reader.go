package file

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/domain"
	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/models"
	"github.com/mauriciozanettisalomao/backseat-driver-kid/pkg/ports"

	"gopkg.in/yaml.v3"
)

type fileInput struct {
	file string
}

func (f *fileInput) Read(ctx context.Context) (*domain.Interaction, error) {

	config, errParseConfig := f.parseConfig(ctx)
	if errParseConfig != nil {
		slog.ErrorContext(ctx, "error parsing config",
			"error", errParseConfig,
			"file", f.file,
		)
		return nil, errParseConfig
	}

	slog.InfoContext(ctx, "config parsed",
		"config", config,
	)

	errParseExtendedKnowledge := f.parseExtendedKnowledge(ctx, config)
	if errParseExtendedKnowledge != nil {
		slog.ErrorContext(ctx, "error parsing extended knowledge",
			"error", errParseExtendedKnowledge,
			"file", f.file,
		)
		return nil, errParseExtendedKnowledge
	}

	errParsePrompt := f.parsePrompts(ctx, config)
	if errParsePrompt != nil {
		slog.ErrorContext(ctx, "error parsing prompts",
			"error", errParsePrompt,
			"file", f.file,
		)
		return nil, errParsePrompt
	}

	return &domain.Interaction{
		Interaction: config.Data.Interaction,
	}, nil

}

func (f *fileInput) parseConfig(ctx context.Context) (*models.ConfigMap, error) {

	inFile, err := os.ReadFile(filepath.Clean(f.file))
	if err != nil {
		slog.ErrorContext(ctx, "error opening file",
			"error", err,
			"file", f.file,
		)
		return nil, err
	}

	var config models.ConfigMap
	err = yaml.Unmarshal(inFile, &config)
	if err != nil {
		log.Fatalf("error unmarshaling YAML: %v", err)
	}

	return &config, nil
}

func (f *fileInput) parsePrompts(ctx context.Context, config *models.ConfigMap) error {

	if config.Data == nil || config.Data.Interaction == nil {
		return errors.New("no prompts found in config")
	}

	if config.Data.Interaction.PromptFile != "" {
		promptsFile, err := os.ReadFile(filepath.Clean(config.Data.Interaction.PromptFile))
		if err != nil {
			slog.ErrorContext(ctx, "error opening prompts file",
				"error", err,
				"file", config.Data.Interaction.PromptFile,
			)
			return err
		}
		for _, prompt := range strings.Split(string(promptsFile), "\n") {
			config.Data.Interaction.Prompts = append(config.Data.Interaction.Prompts, &models.Prompt{
				Input: prompt,
			})
		}
	}

	return nil
}

func (f *fileInput) parseExtendedKnowledge(ctx context.Context, config *models.ConfigMap) error {

	if config.Data == nil || config.Data.Interaction == nil {
		return nil
	}

	if len(config.Data.Interaction.ExtendedKnowledgeDir) > 0 {
		listOfContents := [][]byte{}
		for _, dir := range config.Data.Interaction.ExtendedKnowledgeDir {

			files, errReadDir := os.ReadDir(filepath.Clean(dir))
			if errReadDir != nil {
				slog.ErrorContext(ctx, "error reading extended knowledge dir",
					"error", errReadDir,
					"dir", dir,
				)
				return errReadDir
			}
			for _, file := range files {
				fullpath := fmt.Sprintf("%s/%s", dir, file.Name())
				extendedKnowledgeFile, errReadingExtendedKnowledgeFile := os.ReadFile(filepath.Clean(fullpath))
				if errReadingExtendedKnowledgeFile != nil {
					slog.ErrorContext(ctx, "error opening extended knowledge file",
						"error", errReadingExtendedKnowledgeFile,
						"file", file.Name(),
					)
					return errReadingExtendedKnowledgeFile
				}
				listOfContents = append(listOfContents, extendedKnowledgeFile)
			}
		}
		config.Data.Interaction.ExtendedKnowledgeContent = bytes.Join(listOfContents, []byte("\n\n"))
	}

	return nil
}

// NewInputReader creates a new file input reader
func NewInputReader(file string) ports.InputReader {
	return &fileInput{file}
}
