package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SearchIssueEvents(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchIssueEvents,
) (*mcp.CallToolResult, error) {
	if a.Identifier == "" {
		return response.Fail("identifier is required")
	}

	result, e := s.client.IssueEvents(
		s.organization,
		a.Identifier,
		a.Query,
		a.Limit,
		a.Cursor,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
