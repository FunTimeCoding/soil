package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"sort"
)

func (s *Server) GetSessionTools(
	_ context.Context,
	r server.GetSessionToolsRequestObject,
) (server.GetSessionToolsResponseObject, error) {
	calls := s.service.ToolCalls(r.Identifier)
	var entries []server.ToolCallEntry

	for _, c := range calls {
		entries = append(
			entries,
			server.ToolCallEntry{
				Name:       c.Name,
				Identifier: &c.Identifier,
				Timestamp:  c.Timestamp,
				Detail:     &c.Detail,
			},
		)
	}

	counts := map[string]int{}

	for _, c := range calls {
		counts[c.Name]++
	}

	var sorted []server.ToolCount

	for name, count := range counts {
		sorted = append(
			sorted,
			server.ToolCount{
				Name:  name,
				Count: count,
			},
		)
	}

	sort.Slice(
		sorted,
		func(i, j int) bool {
			return sorted[i].Count > sorted[j].Count
		},
	)

	return server.GetSessionTools200JSONResponse{
		Calls:  entries,
		Counts: sorted,
		Total:  len(calls),
	}, nil
}
