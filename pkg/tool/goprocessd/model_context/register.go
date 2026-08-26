package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.ProcessStatus,
			mcp.WithDescription(
				"List every managed process with its command and whether it is running.",
			),
		),
		mcp.NewTypedToolHandler(s.processStatus),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ProcessLog,
			mcp.WithDescription(
				"Read a process log. Returns the current generation and how many older lines remain.",
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Process name"),
			),
			mcp.WithBoolean(
				"all",
				mcp.Description(
					"Return every retained line instead of the current generation.",
				),
			),
		),
		mcp.NewTypedToolHandler(s.processLog),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ProcessRestart,
			mcp.WithDescription(
				"Restart one process. Use after building a change to a local service.",
			),
			mcp.WithString(
				"name",
				mcp.Required(),
				mcp.Description("Process name"),
			),
		),
		mcp.NewTypedToolHandler(s.processRestart),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ProcessReload,
			mcp.WithDescription(
				"Re-read the Procfile, or re-evaluate the environment file for future restarts.",
			),
			mcp.WithString(
				"scope",
				mcp.Required(),
				mcp.Description("procfile or environment"),
			),
		),
		mcp.NewTypedToolHandler(s.processReload),
	)
}
