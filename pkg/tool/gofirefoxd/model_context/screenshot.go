package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Screenshot(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Screenshot,
) (*mcp.CallToolResult, error) {
	v, e := s.service.Screenshot(int(a.TabIdentifier))

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(v)
}
