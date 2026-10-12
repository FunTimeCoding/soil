package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createVirtualTunnelTermination(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	vmName, e := r.RequireString(constant.VirtualMachine)

	if e != nil {
		return response.Fail("virtual_machine is required: %v", e)
	}

	tunnelName, f := r.RequireString(constant.Tunnel)

	if f != nil {
		return response.Fail("tunnel is required: %v", f)
	}

	interfaceName, g := r.RequireString(constant.Interface)

	if g != nil {
		return response.Fail("interface is required: %v", g)
	}

	role, h := r.RequireString(constant.Role)

	if h != nil {
		return response.Fail("role is required: %v", h)
	}

	tt, i := s.service.CreateVirtualTunnelTermination(
		vmName,
		tunnelName,
		interfaceName,
		role,
	)

	if i != nil {
		return s.serviceFail(i, "tunnel termination not created")
	}

	return response.SuccessAny(convert.TunnelTermination(tt))
}
