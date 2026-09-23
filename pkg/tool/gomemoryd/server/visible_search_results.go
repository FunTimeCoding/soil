package server

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Server) visibleSearchResults(
	results []store.SearchResult,
) []store.SearchResult {
	result := make([]store.SearchResult, 0, len(results))

	for _, r := range results {
		if s.skipHidden(r.Tags) {
			continue
		}

		result = append(result, r)
	}

	return result
}
