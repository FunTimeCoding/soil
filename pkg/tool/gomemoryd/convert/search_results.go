package convert

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func SearchResults(results []store.SearchResult) []*SlimSearchResult {
	result := make([]*SlimSearchResult, 0, len(results))

	for i := range results {
		result = append(result, SearchResult(&results[i]))
	}

	return result
}
