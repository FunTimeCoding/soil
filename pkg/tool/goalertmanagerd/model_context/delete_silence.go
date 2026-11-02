package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) deleteSilence(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeleteSilence,
) (*mcp.CallToolResult, error) {
	if a.Identifier == "" {
		return response.Fail("id is required")
	}

	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	e = s.service.DeleteSilence(x, instance, a.Identifier)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.Success("silence expired: %s", a.Identifier)
}
