package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerAddressing() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListAddresses,
			mcp.WithDescription("List IP addresses assigned to a NetBox device"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListAddresses),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPrefixAddresses,
			mcp.WithDescription(
				"List the IP addresses NetBox holds inside a prefix, each with what holds it: device or virtual machine and interface, role, status, DNS name and description; an address on no interface has neither",
			),
			mcp.WithString(
				constant.Prefix,
				mcp.Required(),
				mcp.Description("Network in CIDR notation (e.g. 192.0.2.0/24)"),
			),
		),
		mcp.NewTypedToolHandler(s.ListPrefixAddresses),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateAddress,
			mcp.WithDescription("Assign an IP address to a device interface"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Interface,
				mcp.Required(),
				mcp.Description("Interface name (e.g. eno1)"),
			),
			mcp.WithString(
				constant.Address,
				mcp.Required(),
				mcp.Description(
					"IP address in CIDR notation (e.g. 203.0.113.60/32)",
				),
			),
			mcp.WithString(
				constant.Status,
				mcp.Description(
					"Address status, e.g. active or reserved (optional)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.CreateAddress),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPrefixes,
			mcp.WithDescription("List all IP prefixes"),
		),
		mcp.NewTypedToolHandler(s.ListPrefixes),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreatePrefix,
			mcp.WithDescription("Create an IP prefix (subnet)"),
			mcp.WithString(
				constant.Prefix,
				mcp.Required(),
				mcp.Description("CIDR notation (e.g. 192.0.2.0/24)"),
			),
			mcp.WithString(
				constant.Site,
				mcp.Description("Site name (optional)"),
			),
			mcp.WithString(
				constant.Description,
				mcp.Description("Description (optional)"),
			),
		),
		mcp.NewTypedToolHandler(s.CreatePrefix),
	)
}
