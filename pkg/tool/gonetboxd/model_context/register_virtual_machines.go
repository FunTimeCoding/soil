package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerVirtualMachines() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListVirtualMachines,
			mcp.WithDescription("List all NetBox virtual machines"),
		),
		s.listVirtualMachines,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateVirtualMachine,
			mcp.WithDescription(
				"Create a NetBox virtual machine assigned to a cluster",
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Cluster,
				mcp.Required(),
				mcp.Description("Cluster name"),
			),
		),
		s.createVirtualMachine,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateVirtualInterface,
			mcp.WithDescription(
				"Create an interface on a NetBox virtual machine",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Interface name (e.g. ens3)"),
			),
		),
		s.createVirtualInterface,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateVirtualAddress,
			mcp.WithDescription(
				"Assign an IP address to a virtual machine interface",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name"),
			),
			mcp.WithString(
				constant.Address,
				mcp.Required(),
				mcp.Description("IP in CIDR notation"),
			),
			mcp.WithString(
				constant.Status,
				mcp.Description(
					"Address status, e.g. active or reserved (optional)",
				),
			),
		),
		s.createVirtualAddress,
	)
}
