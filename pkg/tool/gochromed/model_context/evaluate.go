package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Evaluate(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Evaluate,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Evaluate(
		selector.New(a.TabIdentifier, a.Title, a.Locator),
		a.Expression,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
