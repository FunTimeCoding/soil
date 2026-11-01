package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerDownloads() {
	s.server.AddTool(
		mcp.NewTool(
			constant.DownloadFile,
			mcp.WithDescription(
				"Download a URL through the browser, carrying its cookies and its proxy. Always prefer this over fetching a URL outside the browser, which would bypass Tor. Saves without a prompt and uniquifies a colliding name.",
			),
			mcp.WithString(
				"url",
				mcp.Required(),
				mcp.Description("URL to download"),
			),
			mcp.WithString(
				"filename",
				mcp.Description(
					"Target name, relative to the browser's download directory",
				),
			),
		),
		mcp.NewTypedToolHandler(s.DownloadFile),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DownloadStatus,
			mcp.WithDescription(
				"Report Firefox downloads with their state, bytes received and final path. Omit download_id for the ten most recent.",
			),
			mcp.WithNumber(
				"download_id",
				mcp.Description(
					"Download identifier (omit for the most recent downloads)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.DownloadStatus),
	)
}
