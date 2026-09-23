package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Evaluate(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Evaluate,
) (*mcp.CallToolResult, error) {
	if a.Expression == "" {
		return response.Fail("expression is required")
	}

	t, e := s.resolveTab(a.TabIdentifier, a.Title, a.Locator)

	if e != nil {
		return response.Fail(e.Error())
	}

	p := s.client.Page(t.Identifier)
	var result any
	e = withTimeoutAction(
		constant.TargetTimeout,
		func() error {
			return p.Evaluate(a.Expression, &result)
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
