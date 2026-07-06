package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetIssueTagValues(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetIssueTagValues,
) (*mcp.CallToolResult, error) {
	if a.Identifier == "" {
		return response.Fail("identifier is required")
	}

	if a.Tag == "" {
		return response.Fail("tag is required")
	}

	result, e := s.client.IssueTagValues(
		s.organization,
		a.Identifier,
		a.Tag,
		a.Limit,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
