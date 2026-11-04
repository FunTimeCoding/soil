package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) GetSnippet(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.GetSnippet,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(x, e)
	}

	b, e := s.service.GetSnippet(instance, a.Name)

	if e != nil {
		return s.captureDetail(x, e)
	}

	return response.SuccessAny(
		map[string]string{"name": a.Name, "content": string(b)},
	)
}
