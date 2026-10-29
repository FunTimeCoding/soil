package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) StopPortForward(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.StopPortForward,
) (*mcp.CallToolResult, error) {
	if a.Identifier == "" {
		return response.Fail("id is required")
	}

	if e := s.service.StopPortForward(a.Identifier); e != nil {
		return s.captureDetail(e)
	}

	return response.Success("stopped port forward %s", a.Identifier)
}
