package unit

import (
	textConstant "github.com/funtimecoding/soil/pkg/text/constant"
	"strings"
)

func blocks(pairs ...string) string {
	var lines []string

	for i := 0; i+1 < len(pairs); i += 2 {
		lines = append(lines, textConstant.SearchMarker, pairs[i])
		lines = append(lines, textConstant.DividerMarker)

		if pairs[i+1] != "" {
			lines = append(lines, pairs[i+1])
		}

		lines = append(lines, textConstant.ReplaceMarker)
	}

	return strings.Join(append(lines, ""), "\n")
}
