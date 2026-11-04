package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteJournalEntry(
	x context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireInt(generativeConstant.ParameterIdentifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	if e := s.service.DeleteJournalEntry(int32(identifier)); e != nil {
		return s.captureFail(x, e, "journal entry not deleted")
	}

	return response.Success("journal entry deleted")
}
