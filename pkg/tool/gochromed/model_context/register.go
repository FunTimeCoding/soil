package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
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
			constant.Snapshot,
			mcp.WithDescription(
				"Take an accessibility tree snapshot of a browser tab. Returns structured text with UIDs for each element. Targets the active tab by default. Cross-origin iframe content appears as a bare Iframe node - snapshot the iframe target itself (listed in the result footer and in list_tabs) to read it.",
			),
			mcp.WithString(
				"tab_id",
				mcp.Description("Target tab or iframe target by ID (from list_tabs)"),
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
				mcp.Description("Target tab or iframe target by ID (from list_tabs)"),
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
			constant.Navigate,
			mcp.WithDescription(
				"Navigate a browser tab to a URL. Targets the active tab by default.",
			),
			mcp.WithString(
				"url",
				mcp.Required(),
				mcp.Description("URL to navigate to"),
			),
			mcp.WithString(
				"tab_id",
				mcp.Description("Target tab or iframe target by ID (from list_tabs)"),
			),
		),
		mcp.NewTypedToolHandler(s.Navigate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Evaluate,
			mcp.WithDescription(
				"Run JavaScript in a browser tab and return the result. Targets the active tab by default.",
			),
			mcp.WithString(
				"expression",
				mcp.Required(),
				mcp.Description("JavaScript expression to evaluate"),
			),
			mcp.WithString(
				"tab_id",
				mcp.Description("Target tab by ID"),
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
		mcp.NewTypedToolHandler(s.Evaluate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CloseTab,
			mcp.WithDescription("Close a browser tab."),
			mcp.WithString(
				"tab_id",
				mcp.Description("Target tab by ID"),
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
		mcp.NewTypedToolHandler(s.CloseTab),
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
			constant.ReadBody,
			mcp.WithDescription(
				"Read the HTML body of a browser tab. Saves to download directory and returns the file path. Targets the active tab by default.",
			),
			mcp.WithString(
				"tab_id",
				mcp.Description("Target tab by ID"),
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
		mcp.NewTypedToolHandler(s.ReadBody),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Click,
			mcp.WithDescription(
				"Click an element by UID from the last snapshot. Takes a new snapshot after clicking by default.",
			),
			mcp.WithString(
				"uid",
				mcp.Required(),
				mcp.Description("Element UID from snapshot (e.g. e3)"),
			),
			mcp.WithString(
				"tab_id",
				mcp.Description("Target tab by ID"),
			),
			mcp.WithString(
				"title",
				mcp.Description("Target tab by title substring"),
			),
			mcp.WithString(
				"url",
				mcp.Description("Target tab or iframe target by URL substring"),
			),
			mcp.WithBoolean(
				constant.Snapshot,
				mcp.Description(
					"Take a new snapshot after clicking (default: true)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.Click),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Fill,
			mcp.WithDescription(
				"Fill an input element by UID from the last snapshot. Takes a new snapshot after filling by default.",
			),
			mcp.WithString(
				"uid",
				mcp.Required(),
				mcp.Description("Element UID from snapshot (e.g. e5)"),
			),
			mcp.WithString(
				"value",
				mcp.Required(),
				mcp.Description("Value to fill"),
			),
			mcp.WithString(
				"tab_id",
				mcp.Description("Target tab by ID"),
			),
			mcp.WithString(
				"title",
				mcp.Description("Target tab by title substring"),
			),
			mcp.WithString(
				"url",
				mcp.Description("Target tab or iframe target by URL substring"),
			),
			mcp.WithBoolean(
				constant.Snapshot,
				mcp.Description(
					"Take a new snapshot after filling (default: true)",
				),
			),
			mcp.WithBoolean(
				constant.Direct,
				mcp.Description(
					"Set value directly instead of using CDP Input.insertText (default: false). Use when the default path fails on exotic web components or shadow DOM elements.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.Fill),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Wake,
			mcp.WithDescription(
				"Wake a sleeping browser tab by activating it. Brings the tab to the foreground.",
			),
			mcp.WithString(
				"tab_id",
				mcp.Description("Target tab by ID"),
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
		mcp.NewTypedToolHandler(s.Wake),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.TabHistory,
			mcp.WithDescription(
				"Read the navigation history (back/forward list) of a browser tab. Tab must be awake.",
			),
			mcp.WithString(
				"tab_id",
				mcp.Description("Target tab by ID"),
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
		mcp.NewTypedToolHandler(s.History),
	)
}
