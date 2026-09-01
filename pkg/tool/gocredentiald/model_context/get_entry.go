package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getEntry(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Identifier,
) (*mcp.CallToolResult, error) {
	detail := s.service.Get(a.Identifier)

	if detail == nil {
		return response.Fail("entry not found: %s", a.Identifier)
	}

	return response.SuccessAny(
		map[string]any{
			"identifier":  detail.Identifier,
			"path":        detail.Path,
			"title":       detail.Title,
			"fields":      detail.Fields,
			"modified_at": detail.ModifiedAt,
		},
	)
}
