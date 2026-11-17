package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerLabels() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SetDeviceLabel,
			mcp.WithDescription(
				"Set a key-value label on a device - stored in gonetboxd, not NetBox; upserts by key",
			),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Key,
				mcp.Required(),
				mcp.Description("Label key"),
			),
			mcp.WithString(
				constant.Value,
				mcp.Required(),
				mcp.Description("Label value"),
			),
		),
		mcp.NewTypedToolHandler(s.SetDeviceLabel),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RemoveDeviceLabel,
			mcp.WithDescription("Remove a label from a device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			mcp.WithString(
				constant.Key,
				mcp.Required(),
				mcp.Description("Label key"),
			),
		),
		mcp.NewTypedToolHandler(s.RemoveDeviceLabel),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListDeviceLabels,
			mcp.WithDescription("List labels on a device, ordered by key"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListDeviceLabels),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SetVirtualLabel,
			mcp.WithDescription(
				"Set a key-value label on a virtual machine - stored in gonetboxd, not NetBox; upserts by key",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Key,
				mcp.Required(),
				mcp.Description("Label key"),
			),
			mcp.WithString(
				constant.Value,
				mcp.Required(),
				mcp.Description("Label value"),
			),
		),
		mcp.NewTypedToolHandler(s.SetVirtualLabel),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.RemoveVirtualLabel,
			mcp.WithDescription("Remove a label from a virtual machine"),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
			mcp.WithString(
				constant.Key,
				mcp.Required(),
				mcp.Description("Label key"),
			),
		),
		mcp.NewTypedToolHandler(s.RemoveVirtualLabel),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListVirtualLabels,
			mcp.WithDescription(
				"List labels on a virtual machine, ordered by key",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("VM name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListVirtualLabels),
	)
}
