package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getPageDraft(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(constant.ParameterIdentifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	result, g := s.confluence.DraftOverlay(identifier)

	if g != nil {
		return s.captureFail(g, "page draft not found")
	}

	return response.SuccessAny(result)
}
