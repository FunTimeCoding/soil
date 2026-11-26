package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerActions() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Snapshot,
			mcp.WithDescription(
				"Take a snapshot of a browser tab: an indented tree of roles and names with a UID per element, built from the page's DOM. A snapshot over 20,000 characters is saved to the download directory and its path returned. Targets the first tab by default.",
			),
			mcp.WithString("tab_id", mcp.Description("Target tab by ID")),
		),
		mcp.NewTypedToolHandler(s.Snapshot),
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
				"Fill an input element by UID from the last snapshot, typing the value as a user would. Takes a new snapshot after filling by default.",
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
			mcp.WithBoolean(
				constant.Snapshot,
				mcp.Description(
					"Take a new snapshot after filling (default: true)",
				),
			),
			mcp.WithBoolean(
				constant.Direct,
				mcp.Description(
					"Set the value and fire input and change events instead of typing (default: false). Use when typing does not reach the element.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.Fill),
	)
}
