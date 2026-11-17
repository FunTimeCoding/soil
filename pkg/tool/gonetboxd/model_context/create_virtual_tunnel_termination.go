package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) CreateVirtualTunnelTermination(
	x context.Context,
	_ mcp.CallToolRequest,
	a argument.VirtualTunnelTermination,
) (*mcp.CallToolResult, error) {
	if a.VirtualMachine == "" || a.Tunnel == "" || a.Interface == "" || a.Role == "" {
		return response.Fail(
			"virtual machine, tunnel, interface and role are required",
		)
	}

	tt, e := s.service.CreateVirtualTunnelTermination(
		a.VirtualMachine,
		a.Tunnel,
		a.Interface,
		a.Role,
	)

	if e != nil {
		return s.answerMessage(x, e, "tunnel termination not created")
	}

	return response.SuccessAny(convert.TunnelTermination(tt))
}
