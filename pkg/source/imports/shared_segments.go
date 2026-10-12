package imports

import "strings"

func sharedSegments(
	a string,
	b string,
) int {
	left := strings.Split(a, "/")
	right := strings.Split(b, "/")
	result := 0

	for result < len(left) && result < len(right) {
		if left[result] != right[result] {
			break
		}

		result++
	}

	return result
}
