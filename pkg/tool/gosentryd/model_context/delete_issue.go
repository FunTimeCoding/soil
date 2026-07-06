package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DeleteIssue(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.DeleteIssue,
) (*mcp.CallToolResult, error) {
	if a.Identifier == "" {
		return response.Fail("identifier is required")
	}

	e := s.client.DeleteIssue(s.organization, a.Identifier)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(nil)
}
