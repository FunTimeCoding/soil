package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) list(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	collection, e := q.RequireString(constant.Collection)

	if e != nil {
		return response.Fail("collection is required: %v", e)
	}

	metadata := extractMetadata(q)
	sourceType := q.GetString(constant.SourceType, "")

	if sourceType != "" {
		if metadata == nil {
			metadata = map[string]string{}
		}

		metadata[constant.SourceType] = sourceType
	}

	outcome, f := s.service.ListDocuments(
		collection,
		metadata,
		q.GetInt(parameter.Limit, 10),
		q.GetInt(parameter.Offset, 0),
		q.GetBool(constant.Full, false),
	)

	if f != nil {
		return s.captureFail(f, "document listing failed")
	}

	return response.Success(notation.MarshalIndent(outcome))
}
