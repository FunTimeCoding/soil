package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CloseTab(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.CloseTab,
) (*mcp.CallToolResult, error) {
	if e := s.service.CloseTab(int(a.TabIdentifier)); e != nil {
		return s.answer(e)
	}

	return response.Success("tab closed")
}
