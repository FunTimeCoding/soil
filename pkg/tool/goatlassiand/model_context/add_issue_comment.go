package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addIssueComment(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(generativeConstant.ParameterKey)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	body, g := r.RequireString(generativeConstant.ParameterBody)

	if g != nil {
		return response.Fail("body is required: %v", g)
	}

	if h := s.service.Comment(key, body); h != nil {
		return s.answer(h)
	}

	return response.Success("comment added")
}
