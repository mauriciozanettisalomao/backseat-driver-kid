package formatter

import (
	"bytes"
	"context"
	"html/template"
	"log/slog"
)

func ParseTemplate(ctx context.Context, name, tmpl string, i interface{}) (string, error) {
	t, err := template.New(name).Parse(tmpl)
	if err != nil {
		slog.ErrorContext(ctx, "error parsing template",
			"error", err,
			"name", name,
		)
		return "", err
	}

	var output bytes.Buffer
	if err := t.Execute(&output, i); err != nil {
		slog.ErrorContext(ctx, "error executing template", "error", err)
		return "", err
	}

	return output.String(), nil
}
