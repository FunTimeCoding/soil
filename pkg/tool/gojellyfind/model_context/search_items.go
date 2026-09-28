package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) SearchItems(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.SearchItems,
) (*mcp.CallToolResult, error) {
	items, total, e := s.client.SearchItems(a.Q, a.Types, a.Page, a.PerPage)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]any{"total": total, "page": max(a.Page, 1), "items": items},
	)
}
