package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ReadView(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadView,
) (*mcp.CallToolResult, error) {
	v, e := s.service.View(int(a.ViewIdentifier))

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(v)
}
