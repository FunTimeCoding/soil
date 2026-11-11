package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerSession() {
	s.server.AddTool(
		mcp.NewTool(
			constant.EditSession,
			mcp.WithDescription(
				"Edit session metadata. Set alias, description, topic, or files. All fields optional but at least one required. Defaults to your own session; specify a target to edit another.",
			),
			mcp.WithString(
				constant.Alias,
				mcp.Description(
					"Display name for the session (kebab-case, 2-4 words) - shown in the session list",
				),
			),
			mcp.WithString(
				constant.Description,
				mcp.Description("What the session was about (1-3 sentences)"),
			),
			mcp.WithString(constant.Topic, mcp.Description("Current task")),
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
			constant.Label,
			mcp.WithDescription(
				"Read or change a session label. A label holds a set of values, written comma-separated (role=lead,on_call). Pass value to replace the whole set, add or remove to change one value, nothing to read it.",
			),
			mcp.WithString(
				constant.Key,
				mcp.Required(),
				mcp.Description("Label key"),
			),
			mcp.WithString(
				constant.Value,
				mcp.Description(
					"Replaces the whole set with these comma-separated values; empty string clears the key.",
				),
			),
			mcp.WithString(
				constant.LabelAdd,
				mcp.Description("Adds one value beside the others."),
			),
			mcp.WithString(
				constant.LabelRemove,
				mcp.Description("Removes one value."),
			),
			mcp.WithString(
				constant.Target,
				mcp.Description(
					"Session to label (pool name, alias, or UUID). Omit for own session.",
				),
			),
		),
		s.labelSession,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CruiseTool,
			mcp.WithDescription(
				"Read or set a session's cruise setting - how long it may sit idle before one wake delivers the messages, notifications and pulses waiting for it. Omit grace to read. Setting it on your own session raises your own autonomy and the approval belongs to the person you work with - prefer asking over setting.",
			),
			mcp.WithString(
				constant.Grace,
				mcp.Description(
					fmt.Sprintf(
						"off, default, or idle minutes before the wake: a whole number from %d to %d. Omit to read.",
						constant.CruiseGraceMinimumMinutes,
						constant.CruiseGraceMaximumMinutes,
					),
				),
			),
			mcp.WithString(
				constant.Target,
				mcp.Description(
					"Session (pool name, alias, or UUID). Omit for own session.",
				),
			),
		),
		s.cruiseSession,
	)
}
