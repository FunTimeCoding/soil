package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addCollection(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	name, e := q.RequireString(parameter.Name)

	if e != nil {
		return response.Fail("name is required: %v", e)
	}

	path, f := q.RequireString(constant.Path)

	if f != nil {
		return response.Fail("path is required: %v", f)
	}

	pattern := q.GetString(constant.Pattern, "")
	s.service.AddCollection(name, path, pattern)

	return response.Success("collection %s added", name)
}
