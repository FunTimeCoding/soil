package tool

import "github.com/mark3labs/mcp-go/mcp"

func diskOptions() []mcp.ToolOption {
	return []mcp.ToolOption{
		mcp.WithNumber(
			"disk_size",
			mcp.Description("Disk size in GiB (default 32)"),
		),
		mcp.WithString(
			"disk_storage",
			mcp.Description("Storage backend (default local-lvm)"),
		),
		mcp.WithString(
			"disk_import",
			mcp.Description(
				"Import-from path for disk image, e.g. local:import/debian-13-generic-amd64.qcow2",
			),
		),
		mcp.WithString(
			"cdrom",
			mcp.Description(
				"ISO volume to mount as CD-ROM, e.g. local:iso/debian-13.iso. Boots from CD-ROM first when set.",
			),
		),
	}
}
