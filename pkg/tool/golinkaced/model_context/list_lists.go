package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListLists(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListLists,
) (*mcp.CallToolResult, error) {
	p := pageFromOffset(a.Offset, a.Limit)
	r, e := s.client.ListsPage(p)

	if e != nil {
		return s.captureDetail(e)
	}

	result := make([]map[string]any, len(r.Items))

	for i, l := range r.Items {
		result[i] = map[string]any{
			"identifier": l.Identifier,
			"name":       l.Name,
		}
	}

	result = paginate(result, a.Limit, a.Offset)

	return response.SuccessAny(
		map[string]any{"total": r.Total, "lists": result},
	)
}
