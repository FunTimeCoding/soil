package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"strings"
)

func spanBlock(
	start string,
	closing string,
	end string,
	replace string,
) string {
	lines := []string{constant.SpanMarker, start, closing}

	if end != "" {
		lines = append(lines, end)
	}

	lines = append(lines, constant.DividerMarker)

	if replace != "" {
		lines = append(lines, replace)
	}

	return strings.Join(append(lines, constant.ReplaceMarker, ""), "\n")
}
