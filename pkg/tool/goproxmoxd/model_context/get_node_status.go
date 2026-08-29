package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetNodeStatus(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.GetNodeStatus,
) (*mcp.CallToolResult, error) {
	if a.Node == "" {
		return response.Fail("node is required")
	}

	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	result, e := s.service.GetNodeStatus(instance, a.Node)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(result)
}
