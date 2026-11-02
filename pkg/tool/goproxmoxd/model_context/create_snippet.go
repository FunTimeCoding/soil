package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateSnippet(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.CreateSnippet,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	volume, e := s.service.CreateSnippet(instance, a.Name, a.Content)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(
		map[string]string{"volume": volume, "status": "created"},
	)
}
