package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) updateComment(
	x context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(generativeConstant.ParameterKey)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	identifier, g := r.RequireString(generativeConstant.ParameterIdentifier)

	if g != nil {
		return response.Fail("identifier is required: %v", g)
	}

	body, h := r.RequireString(generativeConstant.ParameterBody)

	if h != nil {
		return response.Fail("body is required: %v", h)
	}

	if e := s.service.UpdateComment(key, identifier, body); e != nil {
		return s.answer(x, e)
	}

	return response.Success("comment %s updated", identifier)
}
