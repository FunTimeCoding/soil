package store

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"

func ExcludePaths(
	results []result.Search,
	exclude []string,
) []result.Search {
	if len(exclude) == 0 {
		return results
	}

	set := map[string]bool{}

	for _, p := range exclude {
		set[p] = true
	}

	var filtered []result.Search

	for _, r := range results {
		if set[r.Path] {
			continue
		}

		filtered = append(filtered, r)
	}

	return filtered
}
