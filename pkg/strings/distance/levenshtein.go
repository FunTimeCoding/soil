package distance

func Levenshtein(
	a string,
	b string,
) int {
	m := len(a)
	n := len(b)

	if m == 0 {
		return n
	}

	if n == 0 {
		return m
	}

	previous := make([]int, n+1)
	current := make([]int, n+1)

	for j := range previous {
		previous[j] = j
	}

	for i := 1; i <= m; i++ {
		current[0] = i

		for j := 1; j <= n; j++ {
			cost := 1

			if a[i-1] == b[j-1] {
				cost = 0
			}

			deletion := previous[j] + 1
			insertion := current[j-1] + 1
			substitution := previous[j-1] + cost
			current[j] = min(deletion, min(insertion, substitution))
		}

		previous, current = current, previous
	}

	return previous[n]
}
