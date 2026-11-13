package unit

import (
	textConstant "github.com/funtimecoding/soil/pkg/text/constant"
	"strings"
)

func spanBlock(
	start string,
	closing string,
	end string,
	replace string,
) string {
	lines := []string{textConstant.SpanMarker, start, closing}

	if end != "" {
		lines = append(lines, end)
	}

	lines = append(lines, textConstant.DividerMarker)

	if replace != "" {
		lines = append(lines, replace)
	}

	return strings.Join(append(lines, textConstant.ReplaceMarker, ""), "\n")
}
