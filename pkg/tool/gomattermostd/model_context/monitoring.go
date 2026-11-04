package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Monitoring(
	x context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	if e := s.service.RunMonitor(); e != nil {
		return s.answer(x, e)
	}

	return response.SuccessAny(
		map[string]any{"message": "Topic monitoring executed successfully"},
	)
}
