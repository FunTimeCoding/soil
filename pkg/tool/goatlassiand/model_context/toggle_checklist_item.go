package model_context

import (
	"context"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) toggleChecklistItem(
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

	items, h := s.service.ToggleChecklistItem(key, int(index))

	if h != nil {
		return s.failOrCapture(h, "checklist not updated")
	}

	return response.SuccessAny(items)
}
