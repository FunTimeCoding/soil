package store

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"

func ExcludePaths(
	results []search.Result,
	exclude []string,
) []search.Result {
	if len(exclude) == 0 {
		return results
	}

	set := map[string]bool{}

	for _, p := range exclude {
		set[p] = true
	}

	var filtered []search.Result

	for _, r := range results {
		if set[r.Path] {
			continue
		}

		filtered = append(filtered, r)
	}

	return filtered
}
