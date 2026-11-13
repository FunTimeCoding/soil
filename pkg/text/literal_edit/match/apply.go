package match

import "strings"

func Apply(
	content string,
	matches []*Match,
) string {
	var result strings.Builder
	previous := 0

	for _, m := range matches {
		result.WriteString(content[previous:m.Offset])
		result.WriteString(m.Replace)
		previous = m.Offset + m.Length
	}

	result.WriteString(content[previous:])

	return result.String()
}
