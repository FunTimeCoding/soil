package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetTabTitle(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SetTabTitle,
) (*mcp.CallToolResult, error) {
	if e := s.service.SetTabTitle(a.TabIdentifier, a.Title); e != nil {
		return s.answer(e)
	}

	return response.Success("set title: %s", a.Title)
}
