package claude_tester

import "fmt"

func ToolResultLine(
	identifier string,
	decision string,
) string {
	return fmt.Sprintf(
		`{"type":"user","permissionDecision":%s,"message":{"content":[{"type":"tool_result","tool_use_id":"%s","content":"done"}]}}`,
		decision,
		identifier,
	)
}
