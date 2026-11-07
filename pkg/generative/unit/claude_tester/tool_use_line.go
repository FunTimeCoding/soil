package claude_tester

import "fmt"

func ToolUseLine(
	identifier string,
	command string,
	directory string,
) string {
	return fmt.Sprintf(
		`{"type":"assistant","cwd":"%s","message":{"content":[{"type":"tool_use","id":"%s","name":"Bash","input":{"command":"%s"}}]}}`,
		directory,
		identifier,
		command,
	)
}
