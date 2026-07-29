package claude

import "github.com/funtimecoding/soil/pkg/generative/constant"

func stripAnsi(s string) string {
	return constant.ClaudeAnsiPattern.ReplaceAllString(s, "")
}
