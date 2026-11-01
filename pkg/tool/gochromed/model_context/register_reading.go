package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerReading() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Snapshot,
			mcp.WithDescription(
				"Take an accessibility tree snapshot of a browser tab. Returns structured text with UIDs for each element. Targets the active tab by default. Cross-origin iframe content appears as a bare Iframe node - snapshot the iframe target itself (listed in the result footer and in list_tabs) to read it.",
			),
			mcp.WithString(
				"tab_id",
				mcp.Description(
					"Target tab or iframe target by ID (from list_tabs)",
				),
			),
			mcp.WithString(
				"title",
				mcp.Description("Target tab by title substring"),
			),
			mcp.WithString(
				"url",
				mcp.Description("Target tab or iframe target by URL substring"),
			),
		),
		mcp.NewTypedToolHandler(s.Snapshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Screenshot,
			mcp.WithDescription(
				"Capture a screenshot of a browser tab. Saves to a temporary file and returns the path for viewing. Targets the active tab by default.",
			),
			mcp.WithString(
				"tab_id",
				mcp.Description(
					"Target tab or iframe target by ID (from list_tabs)",
				),
			),
			mcp.WithString(
				"title",
				mcp.Description("Target tab by title substring"),
			),
			mcp.WithString(
				"url",
				mcp.Description("Target tab or iframe target by URL substring"),
			),
		),
		mcp.NewTypedToolHandler(s.Screenshot),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReadBody,
			mcp.WithDescription(
				"Read the HTML body of a browser tab. Saves to download directory and returns the file path. Targets the active tab by default.",
			),
			mcp.WithString("tab_id", mcp.Description("Target tab by ID")),
			mcp.WithString(
				"title",
				mcp.Description("Target tab by title substring"),
			),
			mcp.WithString(
				"url",
				mcp.Description("Target tab or iframe target by URL substring"),
			),
		),
		mcp.NewTypedToolHandler(s.ReadBody),
	)
}
