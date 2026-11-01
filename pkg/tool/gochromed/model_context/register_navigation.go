package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerNavigation() {
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
				mcp.Description(
					"Target tab or iframe target by ID (from list_tabs)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.Navigate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.TabHistory,
			mcp.WithDescription(
				"Read the navigation history (back/forward list) of a browser tab. Tab must be awake.",
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
		mcp.NewTypedToolHandler(s.History),
	)
}
