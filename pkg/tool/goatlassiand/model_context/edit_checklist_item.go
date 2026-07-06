package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) editChecklistItem(
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

	text, h := r.RequireString(constant.Text)

	if h != nil {
		return response.Fail("text is required: %v", h)
	}

	items, i := s.readChecklist(key)

	if i != nil {
		return s.captureFail(i, "checklist not readable")
	}

	j := int(index)

	if j < 1 || j > len(items) {
		return response.Fail(
			"index %d out of range (1-%d)",
			j,
			len(items),
		)
	}

	items[j-1].Text = text
	fail, e := s.writeChecklist(c, key, items)

	if fail != nil {
		return fail, e
	}

	return response.SuccessAny(items)
}
