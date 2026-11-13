package match

import "strings"

func occurrences(
	content string,
	search string,
) []int {
	var result []int
	start := 0

	for {
		i := strings.Index(content[start:], search)

		if i < 0 {
			return result
		}

		result = append(result, start+i)
		start += i + 1
	}
}
