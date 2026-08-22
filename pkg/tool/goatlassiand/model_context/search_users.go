package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) searchUsers(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, f := r.RequireString(constant.ParameterQuery)

	if f != nil {
		return response.Fail("query is required: %v", f)
	}

	users, g := s.jira.FindUsers(query)

	if g != nil {
		return s.captureDetail(g)
	}

	return response.SuccessAny(convert.JiraUsers(users))
}
