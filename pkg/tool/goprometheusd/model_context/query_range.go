package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/prometheus/client_golang/api/prometheus/v1"
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

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	start, e := parseDurationOrTime(a.Start)

	if e != nil {
		return response.Fail("invalid start: %s", e)
	}

	end := time.Now()

	if a.End != "" {
		end, e = parseDurationOrTime(a.End)

		if e != nil {
			return response.Fail("invalid end: %s", e)
		}
	}

	step := 60 * time.Second

	if a.Step > 0 {
		step = time.Duration(a.Step) * time.Second
	}

	v, f := s.service.QueryRange(
		instance,
		a.Query,
		v1.Range{Start: start, End: end, Step: step},
	)

	if f != nil {
		return s.captureDetail(f)
	}

	return response.SuccessAny(convert.QueryResult(v))
}
