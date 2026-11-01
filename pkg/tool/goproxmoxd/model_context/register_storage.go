package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerStorage() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListStorages,
			mcp.WithDescription(
				"List storage backends on a node. Shows type, content types, and usage.",
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Required(),
				mcp.Description("Node name"),
			),
		),
		mcp.NewTypedToolHandler(s.ListStorages),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListStorageContent,
			mcp.WithDescription(
				"List content on a storage backend. Returns ISOs, disk images, templates, and backups with their volume identifiers.",
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Required(),
				mcp.Description("Node name"),
			),
			mcp.WithString(
				constant.StorageParameter,
				mcp.Required(),
				mcp.Description("Storage name from list_storages"),
			),
		),
		mcp.NewTypedToolHandler(s.ListStorageContent),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DownloadLocator,
			mcp.WithDescription(
				"Download a file from a URL to a storage backend. Use content 'import' for qcow2 cloud images, 'iso' for ISOs, 'vztmpl' for container templates. Waits for the download to complete.",
			),
			mcp.WithString(
				constant.NodeParameter,
				mcp.Required(),
				mcp.Description("Node name"),
			),
			mcp.WithString(
				constant.StorageParameter,
				mcp.Required(),
				mcp.Description("Storage name from list_storages"),
			),
			mcp.WithString(
				constant.ContentParameter,
				mcp.Required(),
				mcp.Description("Content type: import, iso, or vztmpl"),
			),
			mcp.WithString(
				"filename",
				mcp.Required(),
				mcp.Description("Filename to save as on the storage"),
			),
			mcp.WithString(
				"locator",
				mcp.Required(),
				mcp.Description("URL to download from"),
			),
		),
		mcp.NewTypedToolHandler(s.DownloadLocator),
	)
}
