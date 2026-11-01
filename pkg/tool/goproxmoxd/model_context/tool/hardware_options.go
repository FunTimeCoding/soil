package tool

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func hardwareOptions() []mcp.ToolOption {
	return []mcp.ToolOption{
		mcp.WithNumber(
			constant.CoresOption,
			mcp.Description("CPU cores (default 2)"),
		),
		mcp.WithNumber(
			constant.MemoryOption,
			mcp.Description("Memory in MiB (default 2048)"),
		),
		mcp.WithString("cpu_type", mcp.Description("CPU type (default host)")),
		mcp.WithString(
			"os_type",
			mcp.Description("OS type, e.g. l26 for Linux"),
		),
		mcp.WithString(
			"bridge",
			mcp.Description("Network bridge (default vmbr0)"),
		),
	}
}
