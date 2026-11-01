package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerHistory() {
	s.server.AddTool(
		mcp.NewTool(
			constant.History,
			mcp.WithDescription(
				"Show recent events - announcements, completions, messages, moments, summaries. Use to catch up on what happened.",
			),
			mcp.WithNumber(
				constant.Limit,
				mcp.Description("Number of events to show (default 20)"),
			),
			mcp.WithNumber(
				constant.Offset,
				mcp.Description("Number of events to skip (default 0)"),
			),
			mcp.WithString(
				constant.Since,
				mcp.Description("Duration lookback window (e.g. 72h, 168h)"),
			),
			mcp.WithString(
				constant.Before,
				mcp.Description(
					"Duration upper bound (e.g. 72h). Use with since for a time window.",
				),
			),
			mcp.WithString(
				constant.Kind,
				mcp.Description(
					"Filter by event type (e.g. moment, summarize, announce, complete)",
				),
			),
			mcp.WithBoolean(
				constant.Full,
				mcp.Description(
					"Return full event bodies without truncation (default false)",
				),
			),
		),
		s.history,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.HistoryCount,
			mcp.WithDescription(
				"Total number of history events. Quick orientation without fetching content.",
			),
		),
		s.historyCount,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.EditEvent,
			mcp.WithDescription(
				"Edit the body of a history event by ID. Use to fix a hasty summary or correct a milestone description. Editing a summarize or complete event cascades the change to the corresponding record.",
			),
			mcp.WithNumber(
				constant.Identifier,
				mcp.Required(),
				mcp.Description("Event ID from history output (e.g. 42)"),
			),
			mcp.WithString(
				constant.Message,
				mcp.Required(),
				mcp.Description("New body text"),
			),
		),
		s.edit,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Summarize,
			mcp.WithDescription(
				"Record a session summary (tier 2). Captures the full arc of the session - what happened, what was decided, what's next, what landed. One per session, replaces if called again.",
			),
			mcp.WithString(
				constant.Body,
				mcp.Required(),
				mcp.Description("The session summary"),
			),
		),
		s.summarize,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Moment,
			mcp.WithDescription(
				"Capture a line that landed - a moment worth holding. Recorded quietly in the event stream, no notifications.",
			),
			mcp.WithString(
				constant.Line,
				mcp.Required(),
				mcp.Description("The line to capture"),
			),
		),
		s.moment,
	)
}
