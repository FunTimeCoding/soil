package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerRoster() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Roster,
			mcp.WithDescription(
				"List all active sessions with their names and topics.",
			),
		),
		s.roster,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListSessions,
			mcp.WithDescription(
				"Browse all sessions (active and historical) with names, descriptions, and turn counts.",
			),
			mcp.WithNumber(
				constant.Limit,
				mcp.Description("Maximum sessions to return (default 20)"),
			),
			mcp.WithNumber(
				constant.Offset,
				mcp.Description("Number of sessions to skip (default 0)"),
			),
		),
		s.listSessions,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SessionStatus,
			mcp.WithDescription(
				"Show your own session state: name, current topic, files. Use to check what you announced. For the daemon's own problems use status.",
			),
		),
		s.sessionStatus,
	)
}
