package chunk

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"strings"
)

func caption(
	content string,
	lines []Line,
	tableLine int,
) string {
	i := tableLine - 1

	if i >= 0 && trimmed(content, lines[i]) == "" {
		i--
	}

	var result []string

	for ; i >= 0; i-- {
		text := trimmed(content, lines[i])

		if text == "" ||
			constant.AnyHeadingPattern.MatchString(text) ||
			strings.HasPrefix(text, constant.TableRowPrefix) ||
			isFence(text) {
			break
		}

		result = append([]string{text}, result...)
	}

	return strings.Join(result, "\n")
}
