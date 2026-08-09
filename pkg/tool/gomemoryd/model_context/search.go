package model_context

import (
	"context"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) search(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, e := q.RequireString(generative.ParameterQuery)

	if e != nil {
		return response.Fail("query is required")
	}

	limit := int(q.GetFloat(generative.ParameterLimit, 10))
	memoryType := q.GetString(constant.Type, "")
	tag := q.GetString(constant.Tag, "")
	scope := q.GetString(constant.Scope, "")
	results, e := s.service.SearchMemories(
		query,
		limit,
		memoryType,
		tag,
		scope,
	)

	if e != nil {
		return s.captureFail(e, "search failed")
	}

	if q.GetBool(constant.Detail, false) {
		return response.Success(notation.MarshalIndent(results))
	}

	return response.Success(
		notation.MarshalIndent(convert.SearchResults(results)),
	)
}
