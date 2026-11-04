package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) OpenFile(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.OpenFile,
) (*mcp.CallToolResult, error) {
	v, e := s.service.OpenFile(a.FilePath)

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(v)
}
