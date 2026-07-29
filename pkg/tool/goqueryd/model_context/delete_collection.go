package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteCollection(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := q.RequireString(constant.ParameterName)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	if s.service.DeleteCollection(name) {
		return response.Success("collection %s deleted", name)
	}

	return response.Fail("collection not found: %s", name)
}
