package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Navigate(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Navigate,
) (*mcp.CallToolResult, error) {
	if e := s.service.Navigate(a.TabIdentifier, a.Locator); e != nil {
		return s.captureDetail(x, e)
	}

	return response.Success("navigated to %s", a.Locator)
}
