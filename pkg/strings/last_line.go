package strings

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func LastLine(s string) string {
	lines := strings.Split(s, constant.Unix)

	for i := len(lines) - 1; i >= 0; i-- {
		if l := strings.TrimSpace(lines[i]); l != "" {
			return lines[i]
		}
	}

	return ""
}
