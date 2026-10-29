package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Navigate(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Navigate,
) (*mcp.CallToolResult, error) {
	if e := s.service.Navigate(int(a.TabIdentifier), a.Locator); e != nil {
		return s.answer(e)
	}

	return response.Success("navigated to %s", a.Locator)
}
