package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) search(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, e := q.RequireString(parameter.Query)

	if e != nil {
		return response.Fail("query is required")
	}

	limit := int(q.GetFloat(parameter.Limit, 10))
	memoryType := q.GetString(constant.Type, "")
	tag := q.GetString(constant.Tag, "")
	results, e := s.service.SearchMemories(
		query,
		limit,
		memoryType,
		tag,
	)

	if e != nil {
		return s.captureFail(e, "search failed")
	}

	return response.Success(notation.MarshalIndent(results))
}
