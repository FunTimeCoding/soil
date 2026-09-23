package claude_tester

import "fmt"

func AssistantLine(text string) string {
	return fmt.Sprintf(
		`{"type":"assistant","message":{"content":[{"type":"text","text":"%s"}]}}`,
		text,
	)
}
