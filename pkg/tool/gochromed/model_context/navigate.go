package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Navigate(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Navigate,
) (*mcp.CallToolResult, error) {
	if a.Locator == "" {
		return response.Fail("url is required")
	}

	t, e := s.resolveTab(a.TabIdentifier, "", "")

	if e != nil {
		return response.Fail(e.Error())
	}

	p := s.client.Page(t.Identifier)
	e = withTimeoutAction(
		constant.TargetTimeout,
		func() error {
			return p.Navigate(a.Locator)
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("navigated to %s", a.Locator)
}
