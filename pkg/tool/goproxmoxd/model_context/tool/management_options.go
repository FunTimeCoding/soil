package tool

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func managementOptions() []mcp.ToolOption {
	return []mcp.ToolOption{
		mcp.WithBoolean(
			constant.AgentOption,
			mcp.Description("Enable QEMU guest agent (default true)"),
		),
		mcp.WithBoolean(
			constant.OnBootOption,
			mcp.Description("Start VM automatically on host boot"),
		),
		mcp.WithBoolean(
			"start",
			mcp.Description("Start VM after creation (default false)"),
		),
		mcp.WithString(
			constant.TagsOption,
			mcp.Description("Semicolon-separated tags"),
		),
		mcp.WithString(
			"extras",
			mcp.Description(
				"Additional Proxmox options as comma-separated key=value pairs",
			),
		),
	}
}
