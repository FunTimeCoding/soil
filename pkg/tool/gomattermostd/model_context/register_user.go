package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerUser() {
	s.server.AddTool(
		mcp.NewTool(
			constant.MyChannels,
			mcp.WithDescription(
				"List your channels sorted by most recent activity. Includes public, private, DM, and group DM channels.",
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Max channels to return (default 30)"),
				mcp.DefaultNumber(30),
			),
			mcp.WithString(
				"type",
				mcp.Description(
					"Filter by type: D (DM), G (group DM), O (public), P (private). Omit for all.",
				),
			),
			mcp.WithString(
				"since",
				mcp.Description(
					"Only channels with activity since this time (e.g. \"2026-04-22 15:00\")",
				),
			),
		),
		mcp.NewTypedToolHandler(s.MyChannels),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.MyThreads,
			mcp.WithDescription(
				"List your subscribed threads sorted by most recent reply. Shows unread counts, participants, and the original message.",
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Max threads to return (default 20)"),
				mcp.DefaultNumber(20),
			),
			mcp.WithBoolean(
				"unread_only",
				mcp.Description("Only show threads with unread replies"),
			),
			mcp.WithString(
				"since",
				mcp.Description(
					"Only threads with activity since this time (e.g. \"2026-04-22 15:00\")",
				),
			),
		),
		mcp.NewTypedToolHandler(s.MyThreads),
	)
}
