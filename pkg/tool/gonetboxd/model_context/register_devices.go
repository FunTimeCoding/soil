package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerDevices() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListDevices,
			mcp.WithDescription(
				"List NetBox devices, optionally filtered by name",
			),
			mcp.WithString(
				generativeConstant.ParameterQuery,
				mcp.Description("Filter by name (case-insensitive contains)"),
			),
		),
		s.listDevices,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetDevice,
			mcp.WithDescription("Get a NetBox device by exact name"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		s.getDevice,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateDevice,
			mcp.WithDescription(
				"Create a NetBox device. Requires site, role, and device type to exist.",
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Role,
				mcp.Required(),
				mcp.Description("Device role name"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Required(),
				mcp.Description("Device type model name"),
			),
			mcp.WithString(
				constant.Site,
				mcp.Required(),
				mcp.Description("Site name"),
			),
			mcp.WithString(
				constant.Tenant,
				mcp.Description("Tenant name (optional)"),
			),
			mcp.WithArray(
				constant.Tags,
				mcp.WithStringItems(),
				mcp.Description("Tag names to apply (optional)"),
			),
		),
		s.createDevice,
	)
}
