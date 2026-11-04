package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteComment(
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

	if e := s.service.DeleteComment(key, identifier); e != nil {
		return s.answer(x, e)
	}

	return response.Success("comment %s deleted", identifier)
}
