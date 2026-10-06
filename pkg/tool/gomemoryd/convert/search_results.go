package convert

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func SearchResults(results []record.SearchResult) []*SlimSearchResult {
	result := make([]*SlimSearchResult, 0, len(results))

	for i := range results {
		result = append(result, SearchResult(&results[i]))
	}

	return result
}
