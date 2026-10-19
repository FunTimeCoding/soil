package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/convert"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) getStatus(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.GetStatus,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	v, e := s.service.Status(instance)

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(convert.Status(v, a.IncludeConfiguration))
}
