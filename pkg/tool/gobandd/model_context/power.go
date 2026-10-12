package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) power(
	x context.Context,
	operation string,
	force bool,
) (*mcp.CallToolResult, error) {
	instance, e := s.service.ResolveInstance(s.activeInstanceName(x))

	if e != nil {
		return response.Fail("%s", e)
	}

	requested, f := s.service.Power(instance, operation, force)

	if validation.Is(f) || conflict.Is(f) {
		return response.Fail("%s", f)
	}

	if f != nil {
		return s.captureDetail(f)
	}

	return response.Success("%s requested on %s", requested, instance)
}
