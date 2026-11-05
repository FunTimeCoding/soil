package goclaude

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/terminal"
)

func printPeekEntry(
	t *terminal.Terminal,
	userText string,
	assistantContext *string,
	limit int,
) {
	t.Line(userText)

	if limit > 0 && assistantContext != nil && *assistantContext != "" {
		text, _ := cutRunes(*assistantContext, limit)
		t.Line(fmt.Sprintf("  → %s", text))
	}
}
