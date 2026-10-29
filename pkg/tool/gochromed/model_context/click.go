package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium/snapshot"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Click(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Click,
) (*mcp.CallToolResult, error) {
	observe := a.Snapshot == nil || *a.Snapshot
	v, e := s.service.Click(
		selector.New(a.TabIdentifier, a.Title, a.Locator),
		a.UID,
		observe,
	)

	if e != nil {
		return s.captureDetail(e)
	}

	if !observe {
		return response.Success("clicked %s", a.UID)
	}

	return response.Success(
		"clicked %s\n\n%s%s",
		a.UID,
		snapshot.Format(v.Nodes, 0),
		frameHint(v.Frames),
	)
}
