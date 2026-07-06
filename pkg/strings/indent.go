package strings

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
)

func Indent(
	input string,
	count int,
	deleteEmpty bool,
) string {
	lines := split.NewLine(strings.TrimSpace(input))

	if deleteEmpty {
		lines = DeleteEmpty(lines)
	}

	return join.NewLine(IndentLines(lines, count))
}
