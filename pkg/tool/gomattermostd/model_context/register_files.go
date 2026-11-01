package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerFiles() {
	s.server.AddTool(
		mcp.NewTool(
			constant.DownloadFile,
			mcp.WithDescription(
				"Download a file from Mattermost to a local path. Defaults to /tmp/<filename>.",
			),
			mcp.WithString(
				"file_id",
				mcp.Required(),
				mcp.Description("File ID from a post's file_ids"),
			),
			mcp.WithString(
				"path",
				mcp.Description(
					"Local path to save to (default: /tmp/<filename>)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.DownloadFile),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UploadFile,
			mcp.WithDescription(
				"Upload a local file to a Mattermost channel or DM as an attachment.",
			),
			mcp.WithString(
				"path",
				mcp.Required(),
				mcp.Description("Local file path to upload"),
			),
			mcp.WithString("channel_id", mcp.Description("Channel ID")),
			mcp.WithString("channel_name", mcp.Description("Channel name")),
			mcp.WithString(
				"message",
				mcp.Description(
					"Message to post with the file (default: filename)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.UploadFile),
	)
}
