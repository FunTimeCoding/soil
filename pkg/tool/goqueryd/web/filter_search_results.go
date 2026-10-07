package web

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"

func FilterSearchResults(
	outcome *result.Outcome,
	metadata map[string]string,
) ([]result.Search, []result.Facet) {
	if len(metadata) == 0 {
		return outcome.Results, result.ComputeFacets(outcome.Results, 20)
	}

	var filtered []result.Search

	for _, r := range outcome.Results {
		if matchesFilter(r.Metadata, metadata) {
			filtered = append(filtered, r)
		}
	}

	return filtered, result.ComputeFacets(filtered, 20)
}
