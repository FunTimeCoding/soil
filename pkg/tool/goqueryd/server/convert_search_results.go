package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
)

func convertSearchResults(results []result.Search) []server.SearchResult {
	converted := make([]server.SearchResult, len(results))

	for i, r := range results {
		converted[i] = *convertSearchResult(r)
	}

	return converted
}
