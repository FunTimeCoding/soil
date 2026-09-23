package server

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"

func (s *Server) visibleSummaries(
	summaries []store.MemorySummary,
) []store.MemorySummary {
	result := make([]store.MemorySummary, 0, len(summaries))

	for _, m := range summaries {
		if s.skipHidden(m.Tags) {
			continue
		}

		result = append(result, m)
	}

	return result
}
