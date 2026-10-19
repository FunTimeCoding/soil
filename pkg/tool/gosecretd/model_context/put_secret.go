package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) PutSecret(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.PutSecret,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Merge(a.Path, a.Values)

	if e != nil {
		return s.answer(e)
	}

	return response.SuccessAny(map[string]any{"path": a.Path, "keys": result})
}
