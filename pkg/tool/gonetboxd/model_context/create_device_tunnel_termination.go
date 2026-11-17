package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateDeviceTunnelTermination(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.DeviceTunnelTermination,
) (*mcp.CallToolResult, error) {
	if a.Device == "" || a.Tunnel == "" || a.Interface == "" || a.Role == "" {
		return response.Fail("device, tunnel, interface and role are required")
	}

	tt, e := s.service.CreateDeviceTunnelTermination(
		a.Device,
		a.Tunnel,
		a.Interface,
		a.Role,
	)

	if e != nil {
		return s.answerMessage(x, e, "tunnel termination not created")
	}

	return response.SuccessAny(convert.TunnelTermination(tt))
}
