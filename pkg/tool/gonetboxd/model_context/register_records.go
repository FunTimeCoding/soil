package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerRecords() {
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateLocation,
			mcp.WithDescription("Create a NetBox location within a site"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Location name"),
			),
			mcp.WithString(
				constant.Site,
				mcp.Required(),
				mcp.Description("Site name (must already exist)"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateLocation),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListLocations,
			mcp.WithDescription("List all NetBox locations"),
		),
		mcp.NewTypedToolHandler(s.ListLocations),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreatePlatform,
			mcp.WithDescription("Create a NetBox platform"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Platform name"),
			),
		),
		mcp.NewTypedToolHandler(s.CreatePlatform),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPlatforms,
			mcp.WithDescription("List all NetBox platforms"),
		),
		mcp.NewTypedToolHandler(s.ListPlatforms),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateAddressRange,
			mcp.WithDescription("Create a NetBox IP address range"),
			mcp.WithString(
				constant.Start,
				mcp.Required(),
				mcp.Description(
					"Start address in CIDR notation (e.g. 192.168.1.50/24)",
				),
			),
			mcp.WithString(
				constant.End,
				mcp.Required(),
				mcp.Description(
					"End address in CIDR notation (e.g. 192.168.1.254/24)",
				),
			),
			mcp.WithString(
				constant.Status,
				mcp.Description(
					"active, reserved or deprecated; defaults to active",
				),
			),
			mcp.WithString(
				constant.Description,
				mcp.Description("Range description"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateAddressRange),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListAddressRanges,
			mcp.WithDescription("List all NetBox IP address ranges"),
		),
		mcp.NewTypedToolHandler(s.ListAddressRanges),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateWirelessNetwork,
			mcp.WithDescription("Create a NetBox wireless network"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Network name"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateWirelessNetwork),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListWirelessNetworks,
			mcp.WithDescription("List all NetBox wireless networks"),
		),
		mcp.NewTypedToolHandler(s.ListWirelessNetworks),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetChoiceSet,
			mcp.WithDescription("Get a NetBox custom field choice set by name"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Choice set name"),
			),
		),
		mcp.NewTypedToolHandler(s.GetChoiceSet),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListDeviceTags,
			mcp.WithDescription("List the tag names on a NetBox device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListDeviceTags),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteAddress,
			mcp.WithDescription("Delete a NetBox IP address by identifier"),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("IP address identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteAddress),
	)
}
