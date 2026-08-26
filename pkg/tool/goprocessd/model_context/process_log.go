package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) processLog(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ProcessLog,
) (*mcp.CallToolResult, error) {
	lines, older, e := s.supervisor.LogLines(a.Name, a.All)

	if e != nil {
		return response.Fail("%s", e)
	}

	return response.SuccessAny(map[string]any{"lines": lines, "older": older})
}
