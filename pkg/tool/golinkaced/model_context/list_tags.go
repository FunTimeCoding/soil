package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) ListTags(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.ListTags,
) (*mcp.CallToolResult, error) {
	p := pageFromOffset(a.Offset, a.Limit)
	r, e := s.service.TagsPage(p)

	if e != nil {
		return s.answer(x, e)
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
