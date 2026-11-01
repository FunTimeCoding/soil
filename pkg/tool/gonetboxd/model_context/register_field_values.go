package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerFieldValues() {
	field := mcp.WithString(
		constant.Field,
		mcp.Required(),
		mcp.Description("Custom field name"),
	)
	values := mcp.WithArray(
		constant.Values,
		mcp.Required(),
		mcp.WithStringItems(),
		mcp.Description("Selected values; empty clears the field"),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SetDeviceCustomField,
			mcp.WithDescription("Set one custom field on a NetBox device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			field,
			values,
		),
		s.setDeviceCustomField,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SetVirtualCustomField,
			mcp.WithDescription(
				"Set one custom field on a NetBox virtual machine",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("Virtual machine name"),
			),
			field,
			values,
		),
		s.setVirtualCustomField,
	)
}
