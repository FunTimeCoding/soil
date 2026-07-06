package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSessions,
			mcp.WithDescription(
				"List all iTerm2 windows, tabs, and sessions with process context. Returns job_name and command_line for each session, which reveals SSH connections and running programs.",
			),
		),
		mcp.NewTypedToolHandler(s.ListSessions),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReadScreen,
			mcp.WithDescription(
				"Read the visible screen contents of an iTerm2 session. Returns what is currently displayed, including TUI output. Safe for any session type.",
			),
			mcp.WithString(
				"session_id",
				mcp.Required(),
				mcp.Description("iTerm2 session identifier"),
			),
		),
		mcp.NewTypedToolHandler(s.ReadScreen),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ReadHistory,
			mcp.WithDescription(
				"Read scrollback history of an iTerm2 session. Returns recent lines including those scrolled off screen. Useful for shell sessions, meaningless for TUIs.",
			),
			mcp.WithString(
				"session_id",
				mcp.Required(),
				mcp.Description("iTerm2 session identifier"),
			),
			mcp.WithNumber(
				"count",
				mcp.Description("Number of lines to read (default 50)"),
			),
		),
		mcp.NewTypedToolHandler(s.ReadHistory),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SendText,
			mcp.WithDescription(
				"Send text to an iTerm2 session as if typed. Set send_enter to also press enter after the text. Exercise caution: check what session is running (job_name, command_line) before sending to avoid unintended commands on production servers.",
			),
			mcp.WithString(
				"session_id",
				mcp.Required(),
				mcp.Description("iTerm2 session identifier"),
			),
			mcp.WithString(
				"text",
				mcp.Required(),
				mcp.Description("Text to send"),
			),
			mcp.WithBoolean(
				"send_enter",
				mcp.Description(
					"Press enter after sending the text (default false)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.SendText),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SendKey,
			mcp.WithDescription(
				"Send a sequence of named keys to an iTerm2 session with a configurable interval between each. Valid keys: enter, tab, escape, ctrl+c, ctrl+d, ctrl+z, ctrl+l, ctrl+a, ctrl+e, ctrl+r, ctrl+w, ctrl+u, up, down, left, right, backspace, delete. Exercise caution with enter - it executes whatever text is on the command line.",
			),
			mcp.WithString(
				"session_id",
				mcp.Required(),
				mcp.Description("iTerm2 session identifier"),
			),
			mcp.WithArray(
				"keys",
				mcp.Required(),
				mcp.Description(
					"Array of key names to send in sequence (e.g. [\"ctrl+c\", \"up\", \"enter\"])",
				),
				mcp.WithStringItems(),
			),
			mcp.WithNumber(
				"interval",
				mcp.Description(
					"Seconds between each key (default 1, minimum 1)",
				),
			),
		),
		mcp.NewTypedToolHandler(s.SendKey),
	)
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
