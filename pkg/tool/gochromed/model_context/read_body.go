package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ReadBody(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ReadBody,
) (*mcp.CallToolResult, error) {
	v, e := s.service.ReadBody(
		selector.New(a.TabIdentifier, a.Title, a.Locator),
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{"path": v.Path, "tab": v.Title, "size": v.Size},
	)
}
