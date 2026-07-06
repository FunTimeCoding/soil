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

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	var matches []string

	if a.Matches != "" {
		matches = strings.Split(a.Matches, ",")
	}

	since := time.Now().Add(-1 * time.Hour)

	if a.Since != "" {
		d, f := time.ParseDuration(a.Since)

		if f != nil {
			return response.Fail("invalid since: %s", f)
		}

		since = time.Now().Add(-d)
	}

	v, e := s.service.LabelValues(instance, a.Label, matches, since)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(convert.LabelResult(v))
}
