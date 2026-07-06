package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
)

func convertSearchResults(results []store.SearchResult) []server.SearchResult {
	converted := make([]server.SearchResult, len(results))

	for i, r := range results {
		converted[i] = *convertSearchResult(r)
	}

	return converted
}
