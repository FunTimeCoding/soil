package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) processRestartAll(
	_ context.Context,
	_ mcp.CallToolRequest,
	_ struct{},
) (*mcp.CallToolResult, error) {
	count, e := s.service.RestartAll()

	if e != nil {
		return s.answer(e)
	}

	return response.Success(
		"restarting %d processes in the background; watch started_at via process_status",
		count,
	)
}
