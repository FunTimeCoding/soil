package segment

import "strings"

func Segments(name string) []string {
	var result []string

	for _, part := range strings.Split(name, "_") {
		r := []rune(part)
		start := 0

		for i := 1; i < len(r); i++ {
			if !boundary(r, i) {
				continue
			}

			result = append(result, strings.ToLower(string(r[start:i])))
			start = i
		}

		if start < len(r) {
			result = append(result, strings.ToLower(string(r[start:])))
		}
	}

	return result
}
