package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) listMetadata(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	collection, e := q.RequireString(constant.Collection)

	if e != nil {
		return response.Fail("collection is required: %v", e)
	}

	key := q.GetString(constant.Key, "")

	if key != "" {
		return response.Success(
			notation.MarshalIndent(
				s.service.CollectionFacetsForKey(collection, key),
			),
		)
	}

	return response.Success(
		notation.MarshalIndent(
			s.service.CollectionFacets(collection, nil),
		),
	)
}
