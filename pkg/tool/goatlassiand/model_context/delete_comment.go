package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteComment(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(parameter.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	identifier, g := r.RequireString(parameter.Identifier)

	if g != nil {
		return response.Fail("identifier is required: %v", g)
	}

	e := s.jira.Nested().Issue.DeleteCommentWithContext(
		c,
		key,
		identifier,
	)

	if e != nil {
		return response.Fail("delete comment failed: %v", e)
	}

	return response.Success("comment %s deleted", identifier)
}
