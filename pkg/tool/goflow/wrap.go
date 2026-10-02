package goflow

import (
	"strings"
	"unicode/utf8"
)

func wrap(
	units []string,
	width int,
) []string {
	var result []string
	var line strings.Builder
	length := 0

	for _, unit := range units {
		size := utf8.RuneCountInString(unit)

		if length > 0 && length+1+size > width {
			result = append(result, line.String())
			line.Reset()
			length = 0
		}

		if length > 0 {
			line.WriteString(" ")
			length++
		}

		line.WriteString(unit)
		length += size
	}

	if length > 0 {
		result = append(result, line.String())
	}

	return result
}
