package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
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

	v, e := s.client.Evaluate(int(a.TabIdentifier), a.Expression)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(v)
}
