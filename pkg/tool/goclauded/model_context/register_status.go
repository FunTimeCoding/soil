package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerStatus() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Status,
			mcp.WithDescription(
				"Report what is inconsistent inside goclauded: sessions whose transcript is missing, orphaned tracker state, queue entries addressed to a callsign nobody holds, callsigns held past their release window, and callsign pool exhaustion. Use to check whether the coordination state is sound before trusting it, or when something looks off.",
			),
		),
		s.status,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.TokenUsage,
			mcp.WithDescription(
				"Current Claude usage across the five-hour session, seven-day weekly, and Fable limits. Composed from status line reports and the usage page.",
			),
		),
		s.tokenUsage,
	)
}
