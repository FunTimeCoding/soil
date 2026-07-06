package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addContext(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	collection, e := q.RequireString(constant.Collection)

	if e != nil {
		return response.Fail("collection is required: %v", e)
	}

	pathPrefix, f := q.RequireString(constant.PathPrefix)

	if f != nil {
		return response.Fail("path_prefix is required: %v", f)
	}

	description, g := q.RequireString(constant.Description)

	if g != nil {
		return response.Fail("description is required: %v", g)
	}

	s.service.AddContext(collection, pathPrefix, description)

	return response.Success("context added: %s %s", collection, pathPrefix)
}
