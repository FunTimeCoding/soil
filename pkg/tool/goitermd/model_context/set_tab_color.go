package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SetTabColor(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.SetTabColor,
) (*mcp.CallToolResult, error) {
	e := s.service.SetTabColor(
		a.SessionIdentifier,
		int(a.Red),
		int(a.Green),
		int(a.Blue),
	)

	if e != nil {
		return s.answer(x, e)
	}

	return response.Success("color set")
}
