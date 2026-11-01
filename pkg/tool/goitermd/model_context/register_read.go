package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerRead() {
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
}
