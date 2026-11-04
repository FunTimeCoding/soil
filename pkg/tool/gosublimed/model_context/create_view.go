package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateView(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateView,
) (*mcp.CallToolResult, error) {
	v, e := s.service.CreateView(a.Title, a.Content, a.Syntax)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(v)
}
