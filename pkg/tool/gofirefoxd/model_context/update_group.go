package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UpdateGroup(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.UpdateGroup,
) (*mcp.CallToolResult, error) {
	if e := s.service.UpdateGroup(
		int(a.GroupIdentifier),
		a.Title,
		a.Color,
	); e != nil {
		return s.answer(e)
	}

	return response.Success("group updated")
}
