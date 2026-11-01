package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerTabs() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTabs,
			mcp.WithDescription(
				"List all open browser tabs with title, URL, and tab ID. Cross-origin iframes are separate targets and listed after the tabs (type iframe, parent set to the owning tab ID) - all tools accept their IDs and URLs like tabs.",
			),
		),
		mcp.NewTypedToolHandler(s.ListTabs),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTab,
			mcp.WithDescription("Open a new browser tab to a URL."),
			mcp.WithString(
				"url",
				mcp.Required(),
				mcp.Description("URL to open"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateTab),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CloseTab,
			mcp.WithDescription("Close a browser tab."),
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
		mcp.NewTypedToolHandler(s.CloseTab),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Wake,
			mcp.WithDescription(
				"Wake a sleeping browser tab by activating it. Brings the tab to the foreground.",
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
		mcp.NewTypedToolHandler(s.Wake),
	)
}
