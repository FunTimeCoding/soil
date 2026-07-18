package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteJournalEntry(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireInt(parameter.Identifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	if e := s.client.DeleteJournalEntry(int32(identifier)); e != nil {
		return s.captureFail(e, "journal entry not deleted")
	}

	return response.Success("journal entry deleted")
}
