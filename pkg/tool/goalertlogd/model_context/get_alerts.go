package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getAlerts(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, f := r.RequireString(parameter.Name)

	if f != nil {
		return response.Fail("name is required: %v", f)
	}

	records, g := s.store.ByName(name)

	if g != nil {
		return s.captureFail(g, "get alerts failed")
	}

	return response.Success(
		"Found %d alerts:\n%s",
		len(records),
		notation.MarshalIndent(records),
	)
}
