package output

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
)

func FormatConcern(c *concern.Concern) string {
	var location string

	if c.Type == concern.Line && c.Line > 0 {
		location = fmt.Sprintf("%s:%d", c.Path, c.Line)
	} else {
		location = c.Path
	}

	if c.Fixed {
		return fmt.Sprintf("%s: %s (auto-fixed)", location, c.Text)
	}

	return fmt.Sprintf("%s: %s", location, c.Text)
}
