package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Click(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Click,
) (*mcp.CallToolResult, error) {
	v, e := s.service.Click(int(a.TabIdentifier), a.Selector)

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(v)
}
