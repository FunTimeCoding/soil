package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Monitoring(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	s.monitor.RunNow()

	return response.SuccessAny(
		map[string]any{
			"message": "Topic monitoring executed successfully",
		},
	)
}
