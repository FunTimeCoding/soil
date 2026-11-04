package model_context

import (
	"context"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getAlerts(
	x context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(generativeConstant.ParameterName)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	records, g := s.service.Alerts(name)

	if g != nil {
		return s.answer(x, g)
	}

	return response.Success(
		"Found %d alerts:\n%s",
		len(records),
		notation.MarshalIndent(records),
	)
}
