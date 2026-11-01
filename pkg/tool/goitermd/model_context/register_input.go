package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerInput() {
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
}
