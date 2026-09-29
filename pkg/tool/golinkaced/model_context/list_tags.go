package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListTags(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.ListTags,
) (*mcp.CallToolResult, error) {
	p := pageFromOffset(a.Offset, a.Limit)
	r, e := s.client.TagsPage(p)

	if e != nil {
		return s.captureDetail(e)
	}

	result := make([]map[string]any, len(r.Items))

	for i, t := range r.Items {
		result[i] = map[string]any{
			"identifier": t.Identifier,
			"name":       t.Name,
		}
	}

	result = paginate(result, a.Limit, a.Offset)

	return response.SuccessAny(map[string]any{"total": r.Total, "tags": result})
}
