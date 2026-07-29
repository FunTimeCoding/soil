package model_context

import (
	"context"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) updateJournalEntry(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireInt(generative.ParameterIdentifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	comments := r.GetString(constant.Comments, "")
	kind := r.GetString(constant.Kind, "")

	if comments == "" && kind == "" {
		return response.Fail("nothing to update: pass comments or kind")
	}

	result, g := s.client.UpdateJournalEntry(
		int32(identifier),
		kind,
		comments,
	)

	if g != nil {
		return s.captureFail(g, "journal entry not updated")
	}

	return response.SuccessAny(result)
}
