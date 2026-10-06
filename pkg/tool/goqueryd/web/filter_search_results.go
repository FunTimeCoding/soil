package web

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"

func FilterSearchResults(
	outcome *search.Outcome,
	metadata map[string]string,
) ([]search.Result, []search.Facet) {
	if len(metadata) == 0 {
		return outcome.Results, search.ComputeFacets(outcome.Results, 20)
	}

	var filtered []search.Result

	for _, r := range outcome.Results {
		if matchesFilter(r.Metadata, metadata) {
			filtered = append(filtered, r)
		}
	}

	return filtered, search.ComputeFacets(filtered, 20)
}
