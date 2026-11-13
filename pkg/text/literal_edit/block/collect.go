package block

import (
	"slices"
	"strings"
)

func collect(
	lines []string,
	start int,
	markers ...string,
) (string, int, string) {
	for i := start; i < len(lines); i++ {
		if slices.Contains(markers, lines[i]) {
			return strings.Join(lines[start:i], "\n"), i + 1, lines[i]
		}
	}

	return "", len(lines), ""
}
