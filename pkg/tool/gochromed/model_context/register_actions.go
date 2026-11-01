package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerActions() {
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
		mcp.NewTypedToolHandler(s.Evaluate),
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
			mcp.WithString("tab_id", mcp.Description("Target tab by ID")),
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
			mcp.WithString("tab_id", mcp.Description("Target tab by ID")),
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
}
