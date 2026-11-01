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
		s.listAddresses,
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
					"IP address in CIDR notation (e.g. 144.76.220.60/32)",
				),
			),
			mcp.WithString(
				constant.Status,
				mcp.Description(
					"Address status, e.g. active or reserved (optional)",
				),
			),
		),
		s.createAddress,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPrefixes,
			mcp.WithDescription("List all IP prefixes"),
		),
		s.listPrefixes,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreatePrefix,
			mcp.WithDescription("Create an IP prefix (subnet)"),
			mcp.WithString(
				constant.Prefix,
				mcp.Required(),
				mcp.Description("CIDR notation (e.g. 192.168.178.0/24)"),
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
		s.createPrefix,
	)
}
