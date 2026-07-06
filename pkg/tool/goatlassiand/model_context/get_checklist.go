package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getChecklist(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(parameter.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	items, g := s.readChecklist(key)

	if g != nil {
		return s.captureFail(g, "checklist not readable")
	}

	if len(items) == 0 {
		return response.Success("no checklist on %s", key)
	}

	return response.SuccessAny(items)
}
