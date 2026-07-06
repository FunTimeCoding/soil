package claude

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func CleanContent(s string) string {
	s = ansiPattern.ReplaceAllString(s, "")
	s = markupTagPattern.ReplaceAllString(s, " ")
	s = join.Space(strings.Fields(s)...)

	return strings.TrimSpace(s)
}
