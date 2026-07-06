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

	t, e := s.resolveTab(a.TabID, a.Title, a.URL)

	if e != nil {
		return response.Fail(e.Error())
	}

	x := s.client.AcquireTarget(t.Identifier)
	var result any
	e = withTimeoutAction(
		constant.TargetTimeout,
		func() error {
			return s.client.Evaluate(x, a.Expression, &result)
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
