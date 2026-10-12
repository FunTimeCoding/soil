package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Status(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ argument.Status,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Status()

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{
			"initialized": result.Initialized,
			"sealed":      result.Sealed,
			"version":     result.Version,
		},
	)
}
