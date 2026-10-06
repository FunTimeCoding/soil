package chunk

import "strings"

func lineSpans(content string) []line {
	var result []line
	start := 0

	for start < len(content) {
		end := strings.IndexByte(content[start:], '\n')

		if end < 0 {
			result = append(result, line{start: start, end: len(content)})

			break
		}

		result = append(result, line{start: start, end: start + end + 1})
		start += end + 1
	}

	return result
}
