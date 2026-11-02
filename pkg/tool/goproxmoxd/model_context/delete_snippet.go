package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) DeleteSnippet(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeleteSnippet,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	if f := s.service.DeleteSnippet(instance, a.Name); f != nil {
		return s.captureDetail(f)
	}

	return response.Success("deleted")
}
