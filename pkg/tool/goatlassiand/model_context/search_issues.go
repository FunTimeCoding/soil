package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) searchIssues(
	x context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, f := r.RequireString(generativeConstant.ParameterQuery)

	if f != nil {
		return response.Fail("query is required: %v", f)
	}

	limit, g := r.RequireFloat(generativeConstant.ParameterLimit)

	if g != nil {
		return response.Fail("limit is required: %v", g)
	}

	result, h := s.service.SearchIssues(query, new(int(limit)))

	if h != nil {
		return s.answer(x, h)
	}

	return response.SuccessAny(convert.JiraIssues(result))
}
