package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Screenshot(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Screenshot,
) (*mcp.CallToolResult, error) {
	path, e := s.service.Screenshot(a.TabIdentifier)

	if e != nil {
		return s.captureDetail(x, e)
	}

	return response.SuccessAny(map[string]any{"path": path})
}
