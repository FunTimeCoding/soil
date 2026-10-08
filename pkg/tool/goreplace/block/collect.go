package block

import "strings"

func collect(
	lines []string,
	start int,
	marker string,
) (string, int, bool) {
	for i := start; i < len(lines); i++ {
		if lines[i] == marker {
			return strings.Join(lines[start:i], "\n"), i + 1, true
		}
	}

	return "", len(lines), false
}
