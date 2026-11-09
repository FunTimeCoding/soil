package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goprometheusd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
	"time"
)

func (s *Server) labelValues(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.LabelValues,
) (*mcp.CallToolResult, error) {
	if a.Label == "" {
		return response.Fail("label is required")
	}

	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	var matches []string

	if a.Matches != "" {
		matches = strings.Split(a.Matches, ",")
	}

	span := time.Hour

	if a.Since != "" {
		d, f := time.ParseDuration(a.Since)

		if f != nil {
			return response.Fail("invalid since: %s", f)
		}

		span = d
	}

	v, e := s.service.LabelValues(x, instance, a.Label, matches, span)

	if e != nil {
		return s.captureDetail(x, e)
	}

	return response.SuccessAny(convert.LabelResult(v))
}
