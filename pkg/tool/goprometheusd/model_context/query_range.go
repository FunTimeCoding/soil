package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) queryRange(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.QueryRange,
) (*mcp.CallToolResult, error) {
	if a.Query == "" {
		return response.Fail("query is required")
	}

	if a.Start == "" {
		return response.Fail("start is required")
	}

	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	step := 60 * time.Second

	if a.Step > 0 {
		step = time.Duration(a.Step) * time.Second
	}

	v, f := s.service.QueryRange(x, instance, a.Query, a.Start, a.End, step)

	if f != nil {
		return s.captureDetail(x, f)
	}

	return response.SuccessAny(convert.QueryResult(v))
}
