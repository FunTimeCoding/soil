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

	t, e := s.resolveTab(a.TabID, "", "")

	if e != nil {
		return response.Fail(e.Error())
	}

	x := s.client.AcquireTarget(t.Identifier)
	e = withTimeoutAction(
		constant.TargetTimeout,
		func() error {
			return s.client.Navigate(x, a.Locator)
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("navigated to %s", a.Locator)
}
