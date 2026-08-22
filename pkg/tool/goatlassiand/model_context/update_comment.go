package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) updateComment(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(constant.ParameterKey)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	identifier, g := r.RequireString(constant.ParameterIdentifier)

	if g != nil {
		return response.Fail("identifier is required: %v", g)
	}

	body, h := r.RequireString(constant.ParameterBody)

	if h != nil {
		return response.Fail("body is required: %v", h)
	}

	if e := s.service.UpdateComment(key, identifier, body); e != nil {
		return response.Fail("%s", e)
	}

	return response.Success("comment %s updated", identifier)
}
