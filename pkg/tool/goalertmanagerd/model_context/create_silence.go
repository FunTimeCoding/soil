package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func (s *Server) createSilence(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateSilence,
) (*mcp.CallToolResult, error) {
	if a.Alert == "" {
		return response.Fail("alert is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	var d time.Duration

	if a.Duration != "" {
		d, e = time.ParseDuration(a.Duration)

		if e != nil {
			return response.Fail("invalid duration: %s", e)
		}
	}

	v, e := s.service.CreateSilence(instance, a.Alert, a.Comment, d)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("silence created: %s", v)
}
