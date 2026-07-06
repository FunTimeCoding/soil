package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listPages(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	space, f := r.RequireString(constant.SpaceIdentifier)

	if f != nil {
		return response.Fail("space_identifier is required: %v", f)
	}

	status := r.GetString(constant.PageStatus, "")
	result, g := s.confluence.PagesBySpace(space, status)

	if g != nil {
		return s.captureFail(g, "pages not listed")
	}

	return response.SuccessAny(result)
}
