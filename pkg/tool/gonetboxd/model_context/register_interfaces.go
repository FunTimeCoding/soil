package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerInterfaces() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListInterfaces,
			mcp.WithDescription("List interfaces for a NetBox device"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		s.listInterfaces,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateInterface,
			mcp.WithDescription("Create a network interface on a NetBox device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Interface name (e.g. eno1, eth0)"),
			),
			mcp.WithString(
				constant.Type,
				mcp.Required(),
				mcp.Description("Interface type (e.g. 1000base-t, 10gbase-t)"),
			),
			mcp.WithString(
				constant.PhysicalAddress,
				mcp.Description("MAC address to attach (optional)"),
			),
		),
		s.createInterface,
	)
}
