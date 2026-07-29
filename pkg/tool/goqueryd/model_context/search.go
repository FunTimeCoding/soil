package model_context

import (
	"context"
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) search(
	_ context.Context,
	q mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	query, e := q.RequireString(generative.ParameterQuery)

	if e != nil {
		return response.Fail("query is required: %v", e)
	}

	o := search_option.New(query, int(q.GetFloat(generative.ParameterLimit, 10)))
	o.Collection = q.GetString(constant.Collection, "")
	o.Full = q.GetBool(constant.Full, false)
	o.Mode = q.GetString(constant.Mode, "hybrid")
	o.Metadata = extractMetadata(q)
	sourceType := q.GetString(constant.SourceType, "")

	if sourceType != "" {
		if o.Metadata == nil {
			o.Metadata = map[string]string{}
		}

		o.Metadata[constant.SourceType] = sourceType
	}

	outcome := s.service.Search(o)

	if outcome.Cause != nil {
		s.reporter.CaptureException(outcome.Cause)
	}

	return response.Success(notation.MarshalIndent(outcome))
}
