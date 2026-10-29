package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateTab,
) (*mcp.CallToolResult, error) {
	identifier, e := s.service.CreateTab(a.Locator)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{"tab_id": identifier, "url": a.Locator},
	)
}
