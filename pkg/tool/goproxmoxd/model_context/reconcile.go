package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) Reconcile(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.Reconcile,
) (*mcp.CallToolResult, error) {
	instance, e := s.instances.Resolve(x)

	if e != nil {
		return s.captureDetail(e)
	}

	p, e := s.service.Reconcile(x, instance, a.Full, a.Write)

	if unavailable.Is(e) {
		return response.Fail(e.Error())
	}

	if e != nil {
		return s.captureDetail(e)
	}

	return response.SuccessAny(p)
}
