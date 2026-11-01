package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerTunnels() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTunnelGroups,
			mcp.WithDescription("List all VPN tunnel groups"),
		),
		s.listTunnelGroups,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTunnels,
			mcp.WithDescription("List all VPN tunnels"),
		),
		s.listTunnels,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTunnelGroup,
			mcp.WithDescription("Create a VPN tunnel group"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Tunnel group name"),
			),
		),
		s.createTunnelGroup,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTunnel,
			mcp.WithDescription("Create a VPN tunnel in a group"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Tunnel name"),
			),
			mcp.WithString(
				constant.Encapsulation,
				mcp.Required(),
				mcp.Description(
					"Encapsulation type (e.g. wireguard, gre, openvpn)",
				),
			),
			mcp.WithString(
				constant.Group,
				mcp.Required(),
				mcp.Description("Tunnel group name"),
			),
		),
		s.createTunnel,
	)
}
