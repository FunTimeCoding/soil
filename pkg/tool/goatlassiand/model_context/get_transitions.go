package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getTransitions(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(constant.ParameterKey)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	result, g := s.jira.Transitions(key)

	if g != nil {
		return s.captureDetail(g)
	}

	return response.SuccessAny(result)
}
