package server

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Server) visibleSearchResults(
	results []record.SearchResult,
) []record.SearchResult {
	result := make([]record.SearchResult, 0, len(results))

	for _, r := range results {
		if s.skipHidden(r.Tags) {
			continue
		}

		result = append(result, r)
	}

	return result
}
