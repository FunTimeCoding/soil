package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerLifecycle() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Announce,
			mcp.WithDescription(
				"Announce what you're working on. Pass your assigned name (shown in the hook context as 'You are <name>'). Call at the start of a session and when your scope changes.",
			),
			mcp.WithString(
				constant.SessionName,
				mcp.Required(),
				mcp.Description("Your assigned name from the hook context"),
			),
			mcp.WithString(
				constant.Topic,
				mcp.Required(),
				mcp.Description("Short description of what you're working on"),
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
				mcp.Description("Updated files or packages you're touching"),
			),
		),
		s.update,
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
			constant.Release,
			mcp.WithDescription(
				"Release your session from the coordination roster. Call when you're done working.",
			),
		),
		s.release,
	)
}
