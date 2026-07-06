package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteChecklistItem(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(parameter.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	index, g := r.RequireFloat(constant.Index)

	if g != nil {
		return response.Fail("index is required: %v", g)
	}

	items, h := s.readChecklist(key)

	if h != nil {
		return s.captureFail(h, "checklist not readable")
	}

	i := int(index)

	if i < 1 || i > len(items) {
		return response.Fail(
			"index %d out of range (1-%d)",
			i,
			len(items),
		)
	}

	items = append(items[:i-1], items[i:]...)

	for j := range items {
		items[j].Index = j + 1
	}

	fail, e := s.writeChecklist(c, key, items)

	if fail != nil {
		return fail, e
	}

	return response.SuccessAny(items)
}
