package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) RevealSecret(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.RevealSecret,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Reveal(a.Path, a.Key)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{"path": a.Path, "key": a.Key, "value": result},
	)
}
