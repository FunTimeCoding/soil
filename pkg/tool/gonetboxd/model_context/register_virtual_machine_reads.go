package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerVirtualMachineReads() {
	s.server.AddTool(
		mcp.NewTool(
			constant.GetVirtualMachine,
			mcp.WithDescription("Get a NetBox virtual machine with its labels"),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("Virtual machine name"),
			),
		),
		mcp.NewTypedToolHandler(s.GetVirtualMachine),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListVirtualInterfaces,
			mcp.WithDescription("List interfaces on a NetBox virtual machine"),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("Virtual machine name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListVirtualInterfaces),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListVirtualAddresses,
			mcp.WithDescription(
				"List IP addresses assigned to a NetBox virtual machine",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("Virtual machine name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListVirtualAddresses),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListVirtualDisks,
			mcp.WithDescription(
				"List virtual disks on a NetBox virtual machine",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("Virtual machine name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListVirtualDisks),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateVirtualDisk,
			mcp.WithDescription(
				"Create a virtual disk on a NetBox virtual machine",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("Virtual machine name"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Disk name"),
			),
			mcp.WithNumber(
				constant.Size,
				mcp.Required(),
				mcp.Description("Disk size in megabytes"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateVirtualDisk),
	)
}
