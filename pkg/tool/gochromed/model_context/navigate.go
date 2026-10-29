package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Navigate(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Navigate,
) (*mcp.CallToolResult, error) {
	if e := s.service.Navigate(
		selector.New(a.TabIdentifier, "", ""),
		a.Locator,
	); e != nil {
		return s.captureDetail(e)
	}

	return response.Success("navigated to %s", a.Locator)
}
