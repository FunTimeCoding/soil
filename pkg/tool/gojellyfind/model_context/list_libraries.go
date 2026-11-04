package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListLibraries(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.ListLibraries,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Libraries()

	if e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(result)
}
