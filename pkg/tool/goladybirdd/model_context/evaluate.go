package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Evaluate(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Evaluate,
) (*mcp.CallToolResult, error) {
	result, e := s.service.Evaluate(a.TabIdentifier, a.Expression)

	if e != nil {
		return s.captureDetail(x, e)
	}

	return response.Success("%s", result)
}
