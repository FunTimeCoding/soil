package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) status(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.Status,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	v, f := s.service.Status(instance)

	if f != nil {
		return s.captureDetail(f)
	}

	return response.SuccessAny(v)
}
