package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTabs,
			mcp.WithDescription(
				"List the browser's tabs with their IDs and URLs.",
			),
		),
		mcp.NewTypedToolHandler(s.ListTabs),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Navigate,
			mcp.WithDescription(
				"Navigate a browser tab to a URL and wait for it to load. Targets the first tab by default.",
			),
			mcp.WithString(
				"url",
				mcp.Required(),
				mcp.Description("URL to navigate to"),
			),
			mcp.WithString("tab_id", mcp.Description("Target tab by ID")),
		),
		mcp.NewTypedToolHandler(s.Navigate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Evaluate,
			mcp.WithDescription(
				"Evaluate a JavaScript expression in a browser tab, awaiting a returned promise. Returns the result as a WebDriver BiDi remote value. Targets the first tab by default.",
			),
			mcp.WithString(
				"expression",
				mcp.Required(),
				mcp.Description("JavaScript expression to evaluate"),
			),
			mcp.WithString("tab_id", mcp.Description("Target tab by ID")),
		),
		mcp.NewTypedToolHandler(s.Evaluate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Screenshot,
			mcp.WithDescription(
				"Capture a screenshot of a browser tab. Saves a PNG to the download directory and returns its path. Targets the first tab by default.",
			),
			mcp.WithString("tab_id", mcp.Description("Target tab by ID")),
		),
		mcp.NewTypedToolHandler(s.Screenshot),
	)
	s.registerActions()
}
