package claude

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func CleanContent(s string) string {
	s = constant.ClaudeAnsiPattern.ReplaceAllString(s, "")
	s = constant.ClaudeMarkupTagPattern.ReplaceAllString(s, " ")
	s = join.Space(strings.Fields(s)...)

	return strings.TrimSpace(s)
}
