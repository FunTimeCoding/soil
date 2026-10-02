package goflow

import "strings"

func wrap(
	text string,
	width int,
) []string {
	var result []string
	var line strings.Builder

	for _, word := range strings.Fields(text) {
		if line.Len() == 0 {
			line.WriteString(word)

			continue
		}

		if line.Len()+1+len(word) > width {
			result = append(result, line.String())
			line.Reset()
			line.WriteString(word)

			continue
		}

		line.WriteString(" ")
		line.WriteString(word)
	}

	if line.Len() > 0 {
		result = append(result, line.String())
	}

	return result
}
