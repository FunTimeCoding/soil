package server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
)

func (s *Server) GetSessionPeek(
	_ context.Context,
	r server.GetSessionPeekRequestObject,
) (server.GetSessionPeekResponseObject, error) {
	p := s.service.Peek(r.Identifier)
	var entries []server.PeekEntry

	for _, entry := range p.Entries {
		e := server.PeekEntry{
			UserText: entry.UserText,
		}

		if entry.AssistantContext != "" {
			e.AssistantContext = &entry.AssistantContext
		}

		entries = append(entries, e)
	}

	var toolCounts []server.ToolCount

	for _, tc := range p.ToolCounts {
		toolCounts = append(
			toolCounts,
			server.ToolCount{
				Name:  tc.Name,
				Count: tc.Count,
			},
		)
	}

	return server.GetSessionPeek200JSONResponse{
		LineCount:        p.LineCount,
		UserMessageCount: p.UserMessageCount,
		Entries:          entries,
		ToolCounts:       toolCounts,
		TotalToolCalls:   p.TotalToolCalls,
	}, nil
}
