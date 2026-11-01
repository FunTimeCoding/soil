package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerTabs() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListTabs,
			mcp.WithDescription(
				"List all open Firefox tabs with URL, title, and active state",
			),
		),
		mcp.NewTypedToolHandler(s.ListTabs),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTab,
			mcp.WithDescription("Open a URL in a new Firefox tab"),
			mcp.WithString(
				"url",
				mcp.Description("URL to open (defaults to about:blank)"),
			),
		),
		mcp.NewTypedToolHandler(s.CreateTab),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CloseTab,
			mcp.WithDescription("Close a Firefox tab"),
			mcp.WithNumber(
				"tab_id",
				mcp.Required(),
				mcp.Description("Firefox tab identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.CloseTab),
	)
}
