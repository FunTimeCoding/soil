package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) query(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Query,
) (*mcp.CallToolResult, error) {
	if a.Query == "" {
		return response.Fail("query is required")
	}

	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	v, e := s.service.Query(x, instance, a.Query)

	if e != nil {
		return s.captureDetail(x, e)
	}

	return response.SuccessAny(convert.QueryResult(v))
}
