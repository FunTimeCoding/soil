package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium/snapshot"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Snapshot(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Snapshot,
) (*mcp.CallToolResult, error) {
	v, e := s.service.Snapshot(
		selector.New(a.TabIdentifier, a.Title, a.Locator),
	)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success(
		"%s%s",
		snapshot.Format(v.Nodes, 0),
		frameHint(v.Frames),
	)
}
