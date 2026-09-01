package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) loadGroup(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Load,
) (*mcp.CallToolResult, error) {
	group := s.service.LoadGroup(a.Name)

	if group == nil {
		return response.Fail("environment group not found: %s", a.Name)
	}

	return response.SuccessAny(group)
}
