package model_context

import (
	"context"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) editChecklistItem(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(generative.ParameterKey)

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

	items, i := s.service.EditChecklistItem(key, int(index), text)

	if i != nil {
		return s.failOrCapture(i, "checklist not updated")
	}

	return response.SuccessAny(items)
}
