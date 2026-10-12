package unit

import "strings"

func contains(
	s string,
	part string,
) bool {
	return strings.Contains(s, part)
}
