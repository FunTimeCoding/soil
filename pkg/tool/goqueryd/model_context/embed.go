package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) embed(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Embed(x)

	if e != nil {
		return s.answer(x, e)
	}

	return response.Success(
		"embedded %d documents (%d chunks)",
		result.Documents,
		result.Chunks,
	)
}
