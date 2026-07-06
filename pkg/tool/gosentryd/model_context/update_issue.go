package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UpdateIssue(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UpdateIssue,
) (*mcp.CallToolResult, error) {
	if a.Identifier == "" {
		return response.Fail("identifier is required")
	}

	result, e := s.client.UpdateIssue(
		s.organization,
		a.Identifier,
		a.Status,
		a.AssignedTo,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
