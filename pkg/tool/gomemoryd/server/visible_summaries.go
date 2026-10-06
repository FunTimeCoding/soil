package server

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Server) visibleSummaries(
	summaries []record.MemorySummary,
) []record.MemorySummary {
	result := make([]record.MemorySummary, 0, len(summaries))

	for _, m := range summaries {
		if s.skipHidden(m.Tags) {
			continue
		}

		result = append(result, m)
	}

	return result
}
