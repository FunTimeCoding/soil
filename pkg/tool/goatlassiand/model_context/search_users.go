package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) searchUsers(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, f := r.RequireString(parameter.Query)

	if f != nil {
		return response.Fail("query is required: %v", f)
	}

	users, resp, g := s.jira.Nested().User.FindWithContext(
		c,
		query,
	)

	if g != nil {
		if resp != nil && resp.Body != nil {
			return response.Fail(
				"user search failed: %v",
				g,
			)
		}

		return response.Fail("user search failed: %v", g)
	}

	return response.SuccessAny(convert.JiraUsers(users))
}
