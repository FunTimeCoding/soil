package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) search(
	x context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, e := q.RequireString(generativeConstant.ParameterQuery)

	if e != nil {
		return response.Fail("query is required: %v", e)
	}

	o := search_option.New(
		query,
		int(q.GetFloat(generativeConstant.ParameterLimit, 10)),
	)
	o.Collection = q.GetString(constant.Collection, "")
	o.Full = q.GetBool(constant.Full, false)
	o.Mode = q.GetString(constant.Mode, constant.HybridMode)
	o.Metadata = extractMetadata(q)
	sourceType := q.GetString(constant.SourceType, "")

	if sourceType != "" {
		if o.Metadata == nil {
			o.Metadata = map[string]string{}
		}

		o.Metadata[constant.SourceType] = sourceType
	}

	outcome, f := s.service.Search(x, o)

	if unavailable.Is(f) {
		return response.Fail("%v", f)
	}

	if f != nil {
		return s.captureFail(f, "search failed")
	}

	return response.Success(notation.MarshalIndent(outcome))
}
