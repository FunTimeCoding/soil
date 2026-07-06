package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"strings"
)

func formatConcerns(entries []*concern.Concern) string {
	var lines []string

	for _, c := range entries {
		if c.Line > 0 {
			lines = append(
				lines,
				fmt.Sprintf("%s:%d: %s", c.Path, c.Line, c.Text),
			)
		} else if c.Path != "" {
			lines = append(
				lines,
				fmt.Sprintf("%s: %s", c.Path, c.Text),
			)
		} else {
			lines = append(lines, c.Text)
		}
	}

	return strings.Join(lines, "\n")
}
