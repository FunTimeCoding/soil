package model_context

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) updateComment(
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

	body, h := r.RequireString(parameter.Body)

	if h != nil {
		return response.Fail("body is required: %v", h)
	}

	_, resp, i := s.jira.Nested().Issue.UpdateCommentWithContext(
		c,
		key,
		&jira.Comment{ID: identifier, Body: body},
	)

	if i != nil {
		if resp != nil && resp.Body != nil {
			return response.Fail("update comment failed: %v", i)
		}

		return response.Fail("update comment failed: %v", i)
	}

	return response.Success("comment %s updated", identifier)
}
