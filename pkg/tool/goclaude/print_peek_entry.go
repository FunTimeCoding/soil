package goclaude

import "fmt"

func printPeekEntry(
	userText string,
	assistantContext *string,
	limit int,
) {
	fmt.Println(userText)

	if limit > 0 && assistantContext != nil && *assistantContext != "" {
		text := *assistantContext

		if len(text) > limit {
			text = text[:limit]
		}

		fmt.Printf("  → %s\n", text)
	}
}
