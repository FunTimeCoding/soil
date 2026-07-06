package store

func ExcludePaths(
	results []SearchResult,
	exclude []string,
) []SearchResult {
	if len(exclude) == 0 {
		return results
	}

	set := map[string]bool{}

	for _, p := range exclude {
		set[p] = true
	}

	var filtered []SearchResult

	for _, r := range results {
		if set[r.Path] {
			continue
		}

		filtered = append(filtered, r)
	}

	return filtered
}
