package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listGroup(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.Empty,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Groups()

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
