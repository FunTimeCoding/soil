package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listTags(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	tags := s.service.ListSourceTypes()

	if len(tags) == 0 {
		return response.Success("no source type tags configured")
	}

	return response.SuccessAny(tags)
}
