package markup_tester

func heavier(
	a [3]int,
	b [3]int,
) bool {
	for i := range a {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}

	return false
}
