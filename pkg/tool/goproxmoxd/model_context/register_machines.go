package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerMachines() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListMachines,
			mcp.WithDescription(
				"List virtual machines, optionally filtered by node",
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to list across all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListMachines),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetMachine,
			mcp.WithDescription(
				"Get status and config for a virtual machine by VMID",
			),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("VM ID"),
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Description("Node name (omit to search all nodes)"),
			),
		),
		mcp.NewTypedToolHandler(s.GetMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeriveHardwareAddress,
			mcp.WithDescription(
				"Derive the MAC address a machine identifier maps to on the active instance. The identifier need not exist yet - this is how a DHCP reservation is pinned before the machine is created. Reports in_use_by when another guest already carries the address.",
			),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("VM ID, which need not exist yet"),
			),
		),
		mcp.NewTypedToolHandler(s.DeriveHardwareAddress),
	)
}
