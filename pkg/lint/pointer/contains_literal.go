package pointer

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func ContainsLiteral(
	content string,
	needle string,
) bool {
	open := len(needle) > 1 && strings.HasSuffix(needle, constant.Slash)
	needle = strings.TrimSuffix(needle, constant.Slash)

	if open {
		needle = strings.TrimSuffix(needle, constant.Slash)
	}

	offset := 0

	for {
		i := strings.Index(content[offset:], needle)

		if i == -1 {
			return false
		}

		start := offset + i
		end := start + len(needle)
		before := start == 0 || !isPathRune(rune(content[start-1]))
		after := end == len(content) ||
			!isPathRune(rune(content[end])) ||
			(open && content[end] == '/')

		if before && after {
			return true
		}

		offset = start + 1
	}
}
