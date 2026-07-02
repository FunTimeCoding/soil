package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteSilence(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeleteSilence,
) (*mcp.CallToolResult, error) {
	if a.ID == "" {
		return response.Fail("id is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	e = s.service.DeleteSilence(instance, a.ID)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("silence expired: %s", a.ID)
}
