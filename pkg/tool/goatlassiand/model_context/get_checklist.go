package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getChecklist(
	x context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(generativeConstant.ParameterKey)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	items, g := s.service.ReadChecklist(key)

	if g != nil {
		return s.answer(x, g)
	}

	if len(items) == 0 {
		return response.Success("no checklist on %s", key)
	}

	return response.SuccessAny(items)
}
