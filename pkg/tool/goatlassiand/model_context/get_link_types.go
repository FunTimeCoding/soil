package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getLinkTypes(
	_ context.Context,
	_ mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	_, body, e := s.jira.Basic().GetPath(
		"rest/api/2/issueLinkType",
	)

	if e != nil {
		return s.captureDetail(e)
	}

	var parsed linkTypeResponse
	notation.MustDecode(body, &parsed, true)

	return response.SuccessAny(parsed.IssueLinkTypes)
}
