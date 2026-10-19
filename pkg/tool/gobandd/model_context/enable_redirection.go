package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) enableRedirection(
	x context.Context,
	_ mcp.CallToolRequest,
	_ argument.EnableRedirection,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	enabled, f := s.service.EnableRedirection(instance)

	if f != nil {
		return s.captureDetail(f)
	}

	if len(enabled) == 0 {
		return response.Success("redirection already enabled on %s", instance)
	}

	return response.Success(
		"enabled %s on %s",
		join.CommaSpace(enabled),
		instance,
	)
}
