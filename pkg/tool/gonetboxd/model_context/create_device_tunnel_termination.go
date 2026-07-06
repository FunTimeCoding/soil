package model_context

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/mark/response"
	netboxconstant "github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) createDeviceTunnelTermination(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	deviceName, e := r.RequireString(constant.Device)

	if e != nil {
		return response.Fail("device is required: %v", e)
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

	t, i := s.client.TunnelByName(tunnelName)

	if i != nil {
		return s.captureDetail(i)
	}

	d, j := s.client.DeviceByName(deviceName)

	if j != nil {
		return s.captureDetail(j)
	}

	iface, k := s.client.DeviceInterfaceByName(d, interfaceName)

	if k != nil {
		return s.captureDetail(k)
	}

	tt, l := s.client.CreateTunnelTermination(
		t,
		netboxconstant.InterfaceAddress,
		int64(iface.Identifier),
		role,
	)

	if l != nil {
		return s.captureDetail(l)
	}

	return response.SuccessAny(convert.TunnelTermination(tt))
}
