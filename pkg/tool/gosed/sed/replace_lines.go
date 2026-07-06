package sed

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
)

func replaceLines(
	content string,
	replaces map[string]string,
) string {
	lines := split.NewLine(content)

	for i, l := range lines {
		for prefix, replace := range replaces {
			if strings.HasPrefix(strings.TrimSpace(l), prefix) {
				// Keep indentation
				indent := l[:len(l)-len(strings.TrimLeft(l, " \t"))]
				lines[i] = fmt.Sprintf("%s%s%s", indent, prefix, replace)

				break
			}
		}
	}

	return join.NewLine(lines)
}
