package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goreplace/constant"
	"strings"
)

func blocks(pairs ...string) string {
	var lines []string

	for i := 0; i+1 < len(pairs); i += 2 {
		lines = append(lines, constant.SearchMarker, pairs[i])
		lines = append(lines, constant.DividerMarker)

		if pairs[i+1] != "" {
			lines = append(lines, pairs[i+1])
		}

		lines = append(lines, constant.ReplaceMarker)
	}

	return strings.Join(append(lines, ""), "\n")
}
