package strings

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
)

func PrefixMultiline(
	s string,
	prefix string,
) string {
	var lines []string

	for _, line := range split.NewLine(s) {
		lines = append(lines, fmt.Sprintf("%s%s", prefix, line))
	}

	return join.NewLine(lines)
}
