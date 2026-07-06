package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteLink(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(parameter.Identifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	resp, e := s.jira.Nested().Issue.DeleteLinkWithContext(
		c,
		identifier,
	)

	if e != nil {
		if resp != nil && resp.Body != nil {
			return response.Fail("delete link failed: %v", e)
		}

		return response.Fail("delete link failed: %v", e)
	}

	return response.Success("link %s deleted", identifier)
}
