package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerNavigation() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Navigate,
			mcp.WithDescription("Navigate a Firefox tab to a URL"),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
			mcp.WithString(
				"url",
				mcp.Required(),
				mcp.Description("URL to navigate to"),
			),
		),
		mcp.NewTypedToolHandler(s.Navigate),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.NavigateBack,
			mcp.WithDescription("Go back in a Firefox tab's history"),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.NavigateBack),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.WaitFor,
			mcp.WithDescription(
				"Wait for a Firefox tab to finish loading, and optionally until a CSS selector appears. Essential over Tor, where a page takes seconds to settle and a read issued too early returns the previous page. Fails if the selector never appears.",
			),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
			mcp.WithString(
				"selector",
				mcp.Description(
					"CSS selector to wait for (omit to wait only for page load)",
				),
			),
			mcp.WithNumber(
				"timeout",
				mcp.Description("Timeout in milliseconds (default 30000)"),
			),
		),
		mcp.NewTypedToolHandler(s.WaitFor),
	)
}
