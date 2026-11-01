package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerTunnelTerminations() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTunnelTerminations,
			mcp.WithDescription("List all VPN tunnel terminations"),
		),
		s.listTunnelTerminations,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateDeviceTunnelTermination,
			mcp.WithDescription(
				"Create a tunnel termination on a device interface",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Tunnel,
				mcp.Required(),
				mcp.Description("Tunnel name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name on the device"),
			),
			mcp.WithString(
				constant.Role,
				mcp.Required(),
				mcp.Description("Termination role (peer, hub, spoke)"),
			),
		),
		s.createDeviceTunnelTermination,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateVirtualTunnelTermination,
			mcp.WithDescription("Create a tunnel termination on a VM interface"),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Tunnel,
				mcp.Required(),
				mcp.Description("Tunnel name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name on the VM"),
			),
			mcp.WithString(
				constant.Role,
				mcp.Required(),
				mcp.Description("Termination role (peer, hub, spoke)"),
			),
		),
		s.createVirtualTunnelTermination,
	)
}
