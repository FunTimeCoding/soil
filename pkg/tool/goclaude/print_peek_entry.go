package goclaude

import "github.com/funtimecoding/soil/pkg/console"

func printPeekEntry(
	userText string,
	assistantContext *string,
	limit int,
) {
	console.Line(userText)

	if limit > 0 && assistantContext != nil && *assistantContext != "" {
		text := *assistantContext

		if len(text) > limit {
			text = text[:limit]
		}

		console.Format("  → %s\n", text)
	}
}
