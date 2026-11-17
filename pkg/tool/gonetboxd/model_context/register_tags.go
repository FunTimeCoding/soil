package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerTags() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTags,
			mcp.WithDescription("List all NetBox tags"),
		),
		mcp.NewTypedToolHandler(s.ListTags),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTag,
			mcp.WithDescription("Create a NetBox tag"),
			mcp.WithString(
				generativeConstant.ParameterName,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateTag),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddDeviceTag,
			mcp.WithDescription("Add a tag to a device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		mcp.NewTypedToolHandler(s.AddDeviceTag),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RemoveDeviceTag,
			mcp.WithDescription("Remove a tag from a device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		mcp.NewTypedToolHandler(s.RemoveDeviceTag),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AddVirtualTag,
			mcp.WithDescription("Add a tag to a virtual machine"),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		mcp.NewTypedToolHandler(s.AddVirtualTag),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RemoveVirtualTag,
			mcp.WithDescription("Remove a tag from a virtual machine"),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Tag,
				mcp.Required(),
				mcp.Description("Tag name"),
			),
		),
		mcp.NewTypedToolHandler(s.RemoveVirtualTag),
	)
}
