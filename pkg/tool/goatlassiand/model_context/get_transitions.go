package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getTransitions(
	x context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(generativeConstant.ParameterKey)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	result, g := s.service.Transitions(key)

	if g != nil {
		return s.answer(x, g)
	}

	return response.SuccessAny(result)
}
