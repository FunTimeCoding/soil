package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) push(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	collection, e := q.RequireString(constant.Collection)

	if e != nil {
		return response.Fail("collection is required: %v", e)
	}

	path, f := q.RequireString(constant.Path)

	if f != nil {
		return response.Fail("path is required: %v", f)
	}

	body, g := q.RequireString(constant.Body)

	if g != nil {
		return response.Fail("body is required: %v", g)
	}

	metadata := documentMetadata(extractMetadata(q))
	sourceType := q.GetString(constant.SourceType, "")

	if sourceType != "" {
		if metadata == nil {
			metadata = map[string][]string{}
		}

		metadata[constant.SourceType] = []string{sourceType}
	}

	warning := s.service.PushDocument(x, collection, path, body, metadata)

	if warning != "" {
		return response.Success("pushed %s/%s (%s)", collection, path, warning)
	}

	return response.Success("pushed %s/%s", collection, path)
}
