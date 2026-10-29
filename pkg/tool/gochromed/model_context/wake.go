package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Wake(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Wake,
) (*mcp.CallToolResult, error) {
	t, e := s.service.Wake(selector.New(a.TabIdentifier, a.Title, a.Locator))

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("woke tab: %s", t.Title)
}
