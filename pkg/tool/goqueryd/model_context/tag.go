package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) tag(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	path, e := q.RequireString(constant.Path)

	if e != nil {
		return response.Fail("path is required: %v", e)
	}

	collection := q.GetString(constant.Collection, "")
	arguments := q.GetArguments()
	sourceType, hasSourceType := arguments[constant.SourceType]

	if !hasSourceType {
		current := s.service.GetSourceType(collection, path)

		if current == "" {
			return response.Success("no tag for %s", path)
		}

		return response.Success(
			notation.MarshalIndent(
				map[string]string{
					"path":        path,
					"collection":  collection,
					"source_type": current,
				},
			),
		)
	}

	st, _ := sourceType.(string)
	s.service.SetSourceType(collection, path, st)

	if st == "" {
		return response.Success("removed tag for %s", path)
	}

	return response.Success("tagged %s as %s", path, st)
}
