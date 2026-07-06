package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) status(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Status()

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success(notation.MarshalIndent(result))
}
