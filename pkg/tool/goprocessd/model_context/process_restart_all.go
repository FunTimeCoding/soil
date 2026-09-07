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
	count, e := s.supervisor.RestartWave()

	if e != nil {
		return s.captureFail(e, "restart wave failed")
	}

	return response.Success(
		"restarting %d processes in the background; watch started_at via process_status",
		count,
	)
}
