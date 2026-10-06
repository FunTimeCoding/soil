package comparison

import "sort"

func wordDifference(
	from map[string]int,
	to map[string]int,
) ([]string, int) {
	var words []string
	total := 0

	for w, n := range from {
		if d := n - to[w]; d > 0 {
			words = append(words, w)
			total += d
		}
	}

	sort.Strings(words)

	return words, total
}
