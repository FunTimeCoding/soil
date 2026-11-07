package claude_tester

import "fmt"

func RequestBlobLine(identifier string) string {
	return fmt.Sprintf(
		`{"type":"api-request-blob","message":{"content":[{"type":"tool_result","tool_use_id":"%s","content":"done"}]}}`,
		identifier,
	)
}
