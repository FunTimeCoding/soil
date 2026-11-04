package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getLinkTypes(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	types, e := s.service.LinkTypes()

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(types)
}
