package conversations

import "strings"

func nextMatch(
	lower string,
	terms []string,
	position int,
) (int, int) {
	start := -1
	length := 0

	for _, term := range terms {
		i := strings.Index(lower[position:], term)

		if i < 0 {
			continue
		}

		if start < 0 || i < start || i == start && len(term) > length {
			start = i
			length = len(term)
		}
	}

	if start < 0 {
		return -1, 0
	}

	return position + start, length
}
