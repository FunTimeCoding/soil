package model_context

import (
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
			constant.CruiseTool,
			mcp.WithDescription(
				"Read or set a session's cruise setting - whether deliveries sent with cruise may wake it while idle. off (default) waits for the next prompt, on wakes it, paced wakes it at most once per pace. Omit mode to read. Setting it on your own session raises your own autonomy and the approval belongs to the person you work with - prefer asking over setting.",
			),
			mcp.WithString(
				constant.Mode,
				mcp.Enum(
					string(constant.CruiseOff),
					string(constant.CruiseOn),
					string(constant.CruisePaced),
				),
				mcp.Description(
					"Mode to set. Omit to read the current setting.",
				),
			),
			mcp.WithNumber(
				constant.Pace,
				mcp.Description("Minutes between wakes, for paced mode."),
			),
			mcp.WithString(
				constant.Target,
				mcp.Description(
					"Session (pool name, alias, or UUID). Omit for own session.",
				),
			),
		),
		s.cruise,
	)
}
