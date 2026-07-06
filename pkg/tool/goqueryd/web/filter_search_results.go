package web

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/store"

func filterSearchResults(
	outcome *store.SearchOutcome,
	metadata map[string]string,
) ([]store.SearchResult, []store.Facet) {
	if len(metadata) == 0 {
		return outcome.Results, store.ComputeFacets(outcome.Results, 20)
	}

	var filtered []store.SearchResult

	for _, r := range outcome.Results {
		if matchesFilter(r.Metadata, metadata) {
			filtered = append(filtered, r)
		}
	}

	return filtered, store.ComputeFacets(filtered, 20)
}
