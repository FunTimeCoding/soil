package reflow

import (
	markupConstant "github.com/funtimecoding/soil/pkg/markup/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
	"strings"
)

func front(source []byte) int {
	lines := strings.Split(string(source), "\n")

	if len(lines) == 0 || lines[0] != markupConstant.FrontMatterDelimiter {
		return 0
	}

	at := slices.Index(lines[1:], markupConstant.FrontMatterDelimiter)

	if at < 0 {
		return 0
	}

	return len(join.NewLine(lines[:at+2])) + 1
}
