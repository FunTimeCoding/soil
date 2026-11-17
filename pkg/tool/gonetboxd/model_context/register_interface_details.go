package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerInterfaceDetails() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SetInterfacePhysicalAddress,
			mcp.WithDescription(
				"Assign a MAC address to an existing device interface",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name"),
			),
			mcp.WithString(
				constant.Address,
				mcp.Required(),
				mcp.Description("MAC address"),
			),
		),
		mcp.NewTypedToolHandler(s.SetInterfacePhysicalAddress),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SetVirtualInterfacePhysicalAddress,
			mcp.WithDescription(
				"Assign a MAC address to an existing virtual machine interface",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("Virtual machine name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name"),
			),
			mcp.WithString(
				constant.Address,
				mcp.Required(),
				mcp.Description("MAC address"),
			),
		),
		mcp.NewTypedToolHandler(s.SetVirtualInterfacePhysicalAddress),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SetInterfaceEthernetPower,
			mcp.WithDescription(
				"Set a device interface's ethernet power - the NetBox PoE fields",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name"),
			),
			mcp.WithString(
				constant.Mode,
				mcp.Description(
					"pse when the port supplies power, pd when it is powered; left out or empty clears mode and type",
				),
			),
			mcp.WithString(
				constant.Type,
				mcp.Description("The standard, such as type2-ieee802.3at"),
			),
		),
		mcp.NewTypedToolHandler(s.SetInterfaceEthernetPower),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPhysicalAddresses,
			mcp.WithDescription(
				"List every MAC address and the device or virtual machine carrying it",
			),
		),
		mcp.NewTypedToolHandler(s.ListPhysicalAddresses),
	)
}
