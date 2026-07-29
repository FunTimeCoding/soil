package model_context

import (
	"context"
	"github.com/andygrunwald/go-jira"
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
