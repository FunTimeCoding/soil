package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListLinks(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListLinks,
) (*mcp.CallToolResult, error) {
	p := pageFromOffset(a.Offset, a.Limit)

	if a.List > 0 {
		r, e := s.client.LinksByListPage(a.List, p)

		if e != nil {
			return s.captureDetail(e)
		}

		return s.linkPage(r.Items, r.Total, a.Limit, a.Offset)
	}

	r, e := s.client.LinksPage(p)

	if e != nil {
		return s.captureDetail(e)
	}

	return s.linkPage(r.Items, r.Total, a.Limit, a.Offset)
}
