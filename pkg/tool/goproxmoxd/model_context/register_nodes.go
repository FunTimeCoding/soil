package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerNodes() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListNodes,
			mcp.WithDescription("List all Proxmox nodes with status"),
		),
		mcp.NewTypedToolHandler(s.ListNodes),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetNodeStatus,
			mcp.WithDescription(
				"Get detailed resource status for a Proxmox node",
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Required(),
				mcp.Description("Node name"),
			),
		),
		mcp.NewTypedToolHandler(s.GetNodeStatus),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListNetworks,
			mcp.WithDescription("List network interfaces on a Proxmox node"),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Required(),
				mcp.Description("Node name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListNetworks),
	)
}
