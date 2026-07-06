package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) searchIssues(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, f := r.RequireString(parameter.Query)

	if f != nil {
		return response.Fail("query is required: %v", f)
	}

	limit, g := r.RequireFloat(parameter.Limit)

	if g != nil {
		return response.Fail("limit is required: %v", g)
	}

	result, h := s.jira.SearchLimit(int(limit), query)

	if h != nil {
		return s.captureDetail(h)
	}

	return response.SuccessAny(convert.JiraIssues(result))
}
