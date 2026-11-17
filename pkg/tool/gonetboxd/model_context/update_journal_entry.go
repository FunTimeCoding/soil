package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UpdateJournalEntry(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.JournalEntryChange,
) (*mcp.CallToolResult, error) {
	if a.Comments == "" && a.Kind == "" {
		return response.Fail("nothing to update: pass comments or kind")
	}

	result, e := s.service.UpdateJournalEntry(
		int32(a.Identifier),
		a.Kind,
		a.Comments,
	)

	if e != nil {
		return s.answerMessage(x, e, "journal entry not updated")
	}

	return response.SuccessAny(result)
}
