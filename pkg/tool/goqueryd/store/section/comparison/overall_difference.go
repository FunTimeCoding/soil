package comparison

func overallDifference(
	from map[string]map[string]int,
	to map[string]map[string]int,
) (int, int) {
	balance := map[string]int{}

	for _, m := range from {
		for w, n := range m {
			balance[w] += n
		}
	}

	for _, m := range to {
		for w, n := range m {
			balance[w] -= n
		}
	}

	removed, added := 0, 0

	for _, d := range balance {
		if d > 0 {
			removed += d
		} else {
			added -= d
		}
	}

	return removed, added
}
