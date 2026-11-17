package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DeleteJournalEntry(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Identifier,
) (*mcp.CallToolResult, error) {
	if e := s.service.DeleteJournalEntry(int32(a.Identifier)); e != nil {
		return s.answerMessage(x, e, "journal entry not deleted")
	}

	return response.Success("journal entry deleted")
}
