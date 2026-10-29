package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GroupTabs(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GroupTabs,
) (*mcp.CallToolResult, error) {
	groupIdentifier, e := s.service.GroupTabs(
		a.TabIdentifiers,
		int(a.GroupIdentifier),
		a.Title,
		a.Color,
	)

	if e != nil {
		return s.answer(e)
	}

	return response.Success("grouped into group %d", groupIdentifier)
}
