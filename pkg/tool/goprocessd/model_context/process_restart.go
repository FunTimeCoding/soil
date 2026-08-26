package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) processRestart(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ProcessRestart,
) (*mcp.CallToolResult, error) {
	if e := s.supervisor.Restart([]string{a.Name}); e != nil {
		return s.captureFail(e, "restart failed")
	}

	return response.SuccessAny(s.supervisor.Statuses())
}
