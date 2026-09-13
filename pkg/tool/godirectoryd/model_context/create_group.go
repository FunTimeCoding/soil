package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createGroup(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Group,
) (*mcp.CallToolResult, error) {
	result, e := s.service.CreateGroup(a.Name)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
