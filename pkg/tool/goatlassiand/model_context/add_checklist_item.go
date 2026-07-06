package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addChecklistItem(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(parameter.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	text, g := r.RequireString(constant.Text)

	if g != nil {
		return response.Fail("text is required: %v", g)
	}

	items, h := s.readChecklist(key)

	if h != nil {
		return s.captureFail(h, "checklist not readable")
	}

	items = append(
		items,
		checklist_item.New(len(items), text, false),
	)

	if fail, e := s.writeChecklist(c, key, items); fail != nil {
		return fail, e
	}

	return response.SuccessAny(items)
}
