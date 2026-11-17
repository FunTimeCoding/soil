package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerPower() {
	s.server.AddTool(
		mcp.NewTool(
			constant.CreatePowerOutlet,
			mcp.WithDescription(
				"Create a power outlet on a device, such as a UPS or a power strip",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Outlet name, such as battery 1"),
			),
		),
		mcp.NewTypedToolHandler(s.CreatePowerOutlet),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreatePowerPort,
			mcp.WithDescription(
				"Create a power port, the inlet a device is powered through",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Port name, such as power"),
			),
		),
		mcp.NewTypedToolHandler(s.CreatePowerPort),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreatePowerCable,
			mcp.WithDescription(
				"Create a power cable from a power outlet to a power port",
			),
			mcp.WithString(
				constant.OutletDevice,
				mcp.Required(),
				mcp.Description("Device carrying the outlet"),
			),
			mcp.WithString(
				constant.Outlet,
				mcp.Required(),
				mcp.Description("Power outlet name"),
			),
			mcp.WithString(
				constant.PortDevice,
				mcp.Required(),
				mcp.Description("Device whose power port the cable plugs into"),
			),
			mcp.WithString(
				constant.Port,
				mcp.Required(),
				mcp.Description("Power port name"),
			),
		),
		mcp.NewTypedToolHandler(s.CreatePowerCable),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListPowerOutlets,
			mcp.WithDescription(
				"List a device's power outlets and what each one powers",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListPowerOutlets),
	)
}
