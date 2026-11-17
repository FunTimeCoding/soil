package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerCabling() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListCables,
			mcp.WithDescription("List all NetBox cables"),
		),
		mcp.NewTypedToolHandler(s.ListCables),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateCable,
			mcp.WithDescription(
				"Create a NetBox cable between two device interfaces",
			),
			mcp.WithString(
				constant.DeviceA,
				mcp.Required(),
				mcp.Description("Device name on side A"),
			),
			mcp.WithString(
				constant.InterfaceA,
				mcp.Required(),
				mcp.Description("Interface name on side A"),
			),
			mcp.WithString(
				constant.DeviceB,
				mcp.Required(),
				mcp.Description("Device name on side B"),
			),
			mcp.WithString(
				constant.InterfaceB,
				mcp.Required(),
				mcp.Description("Interface name on side B"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateCable),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteCable,
			mcp.WithDescription(
				"Delete a NetBox cable by identifier, as list_cables shows it",
			),
			mcp.WithNumber(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Cable identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteCable),
	)
}
