package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getLinkTypes(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	types, e := s.service.LinkTypes()

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(types)
}
