package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addPageComment(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(parameter.Identifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	body, g := r.RequireString(parameter.Body)

	if g != nil {
		return response.Fail("body is required: %v", g)
	}

	if h := s.confluence.AddComment(identifier, body); h != nil {
		return s.captureFail(h, "comment not added")
	}

	return response.Success("comment added")
}
