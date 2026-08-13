package status

import "fmt"

func extendLines(
	input string,
	lines []string,
	indent int,
) string {
	result := input

	for _, line := range lines {
		result = fmt.Sprintf("%s\n%s%s", result, spaces(indent), line)
	}

	return result
}
