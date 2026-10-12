package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListSecrets(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListSecrets,
) (*mcp.CallToolResult, error) {
	result, e := s.service.List(a.Path)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{"path": a.Path, "entries": result},
	)
}
