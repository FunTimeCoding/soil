package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerTabs() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SetTabTitle,
			mcp.WithDescription("Set the title of an iTerm2 tab"),
			mcp.WithString(
				"tab_id",
				mcp.Required(),
				mcp.Description("Tab identifier from list_sessions"),
			),
			mcp.WithString(
				"title",
				mcp.Required(),
				mcp.Description("New tab title"),
			),
		),
		mcp.NewTypedToolHandler(s.SetTabTitle),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SetTabColor,
			mcp.WithDescription("Set the color of an iTerm2 tab"),
			mcp.WithString(
				"session_id",
				mcp.Required(),
				mcp.Description("iTerm2 session identifier"),
			),
			mcp.WithNumber(
				"red",
				mcp.Required(),
				mcp.Description("Red component (0-255)"),
			),
			mcp.WithNumber(
				"green",
				mcp.Required(),
				mcp.Description("Green component (0-255)"),
			),
			mcp.WithNumber(
				"blue",
				mcp.Required(),
				mcp.Description("Blue component (0-255)"),
			),
		),
		mcp.NewTypedToolHandler(s.SetTabColor),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateTab,
			mcp.WithDescription("Create a new tab in the current iTerm2 window"),
		),
		mcp.NewTypedToolHandler(s.CreateTab),
	)
}
