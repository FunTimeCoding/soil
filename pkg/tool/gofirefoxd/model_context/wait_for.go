package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) WaitFor(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.WaitFor,
) (*mcp.CallToolResult, error) {
	v, e := s.service.WaitFor(int(a.TabIdentifier), a.Selector, int(a.Timeout))

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(v)
}
