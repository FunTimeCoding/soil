package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListViews(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListViews,
) (*mcp.CallToolResult, error) {
	v, e := s.service.Views()

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(v)
}
