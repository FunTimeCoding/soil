package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/chromium/snapshot"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gochromed/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gochromed/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Snapshot(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.Snapshot,
) (*mcp.CallToolResult, error) {
	t, e := s.resolveTab(a.TabID, a.Title, a.URL)

	if e != nil {
		return response.Fail(e.Error())
	}

	x := s.client.AcquireTarget(t.Identifier)
	nodes, e := withTimeout(
		constant.TargetTimeout,
		func() ([]*snapshot.Node, error) {
			return s.client.Snapshot(x)
		},
	)

	if e != nil {
		return s.captureDetail(e)
	}

	s.cacheSnapshot(t.Identifier, nodes)

	return response.Success(
		"%s%s",
		snapshot.Format(nodes, 0),
		s.frameHint(t.Identifier),
	)
}
