package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) UngroupTab(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.UngroupTab,
) (*mcp.CallToolResult, error) {
	if e := s.service.UngroupTab(int(a.TabIdentifier)); e != nil {
		return s.answer(x, e)
	}

	return response.Success("ungrouped")
}
