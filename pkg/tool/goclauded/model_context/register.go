package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Announce,
			mcp.WithDescription(
				"Announce what you're working on. Pass your assigned name (shown in the hook context as 'You are <name>'). Call at the start of a session and when your scope changes.",
			),
			mcp.WithString(
				constant.SessionName,
				mcp.Required(),
				mcp.Description(
					"Your assigned name from the hook context",
				),
			),
			mcp.WithString(
				constant.Topic,
				mcp.Required(),
				mcp.Description(
					"Short description of what you're working on",
				),
			),
			mcp.WithArray(
				constant.Files,
				mcp.Description("Files or packages you're touching"),
			),
		),
		s.announce,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Complete,
			mcp.WithDescription(
				"Mark the current task complete. Clears your topic so you can announce the next task. Provide a short summary of what was accomplished. Calling again on the same topic amends the existing completion.",
			),
			mcp.WithString(
				constant.Message,
				mcp.Required(),
				mcp.Description(
					"Summary of what was accomplished (1-2 sentences)",
				),
			),
			mcp.WithString(
				constant.Topic,
				mcp.Description(
					"Topic to complete, if no active topic was announced",
				),
			),
		),
		s.complete,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Update,
			mcp.WithDescription(
				"Record a milestone and shift to the next scope. Captures what was accomplished before moving on. Use between announce and complete to track progress incrementally.",
			),
			mcp.WithString(
				constant.Message,
				mcp.Required(),
				mcp.Description("What was accomplished in the current scope"),
			),
			mcp.WithString(
				constant.Topic,
				mcp.Required(),
				mcp.Description("New topic description"),
			),
			mcp.WithArray(
				constant.Files,
				mcp.Description(
					"Updated files or packages you're touching",
				),
			),
		),
		s.update,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Status,
			mcp.WithDescription(
				"Show your own session state: name, current topic, files. Use to check what you announced.",
			),
		),
		s.status,
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
				mcp.Description(
					"Event ID from history output (e.g. 42)",
				),
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
				mcp.Description(
					"Maximum sessions to return (default 20)",
				),
			),
			mcp.WithNumber(
				constant.Offset,
				mcp.Description(
					"Number of sessions to skip (default 0)",
				),
			),
		),
		s.listSessions,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.History,
			mcp.WithDescription(
				"Show recent events - announcements, completions, messages, moments, summaries. Use to catch up on what happened.",
			),
			mcp.WithNumber(
				constant.Limit,
				mcp.Description(
					"Number of events to show (default 20)",
				),
			),
			mcp.WithNumber(
				constant.Offset,
				mcp.Description("Number of events to skip (default 0)"),
			),
			mcp.WithString(
				constant.Since,
				mcp.Description(
					"Duration lookback window (e.g. 72h, 168h)",
				),
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
			constant.EditSession,
			mcp.WithDescription(
				"Edit session metadata. Set alias, description, topic, or files. All fields optional but at least one required. Defaults to your own session; specify a target to edit another.",
			),
			mcp.WithString(
				constant.Alias,
				mcp.Description("Display name for the session"),
			),
			mcp.WithString(
				constant.Description,
				mcp.Description(
					"What the session was about (1-3 sentences)",
				),
			),
			mcp.WithString(
				constant.Topic,
				mcp.Description("Current task"),
			),
			mcp.WithString(
				constant.Files,
				mcp.Description("Files or packages being touched"),
			),
			mcp.WithString(
				constant.Target,
				mcp.Description(
					"Session to edit (pool name, alias, or UUID). Omit to edit yourself.",
				),
			),
		),
		s.editSession,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Send,
			mcp.WithDescription(
				"Send a message to another session by name, or broadcast to all. Use for coordination: build warnings, scope handoffs, questions.",
			),
			mcp.WithString(
				constant.To,
				mcp.Description(
					"Recipient name (omit to broadcast to all)",
				),
			),
			mcp.WithString(
				constant.Body,
				mcp.Required(),
				mcp.Description("Message text"),
			),
		),
		s.send,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Listen,
			mcp.WithDescription(
				"Opt in to being woken by messages from other sessions. When listening, inbound messages will wake you in a new turn.",
			),
		),
		s.listen,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Release,
			mcp.WithDescription(
				"Release your session from the coordination roster. Call when you're done working.",
			),
		),
		s.release,
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
	s.server.AddTool(
		mcp.NewTool(
			constant.Label,
			mcp.WithDescription(
				"Set, get, or remove a session label. Pass value to set, empty string to remove, omit to get current value.",
			),
			mcp.WithString(
				constant.Key,
				mcp.Required(),
				mcp.Description("Label key"),
			),
			mcp.WithString(
				constant.Value,
				mcp.Description(
					"Label value. Pass to set, empty string to remove, omit to get.",
				),
			),
			mcp.WithString(
				constant.Target,
				mcp.Description(
					"Session to label (pool name, alias, or UUID). Omit for own session.",
				),
			),
		),
		s.label,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Pulse,
			mcp.WithDescription(
				"Send a pulse to the dashboard. Appears on the session card and detail page. Bidirectional - the dashboard can pulse back.",
			),
			mcp.WithString(
				constant.Line,
				mcp.Required(),
				mcp.Description("The line to send"),
			),
		),
		s.pulse,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.TokenUsage,
			mcp.WithDescription(
				"Current Claude usage across session, weekly, and credit limits. Read from the claude.ai usage page.",
			),
		),
		s.tokenUsage,
	)
}
