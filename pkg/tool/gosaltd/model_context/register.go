package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Sync,
			mcp.WithDescription(
				"Pull the latest changes from the repository. Returns a filtered log of commits that touched salt states.",
			),
		),
		s.syncRepository,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Trigger,
			mcp.WithDescription(
				"Trigger a salt highstate. Set target to run on specific minions. Set update to pull latest code first. Set synchronous to wait for the result instead of returning immediately.",
			),
			mcp.WithString(
				constant.Target,
				mcp.Description(
					"Target minion pattern. Default: * (all minions).",
				),
			),
			mcp.WithBoolean(
				constant.Update,
				mcp.Description(
					"Pull latest code before running. Default: false.",
				),
			),
			mcp.WithBoolean(
				constant.Synchronous,
				mcp.Description(
					"Wait for the run to complete and return the result. Default: false.",
				),
			),
		),
		s.triggerRun,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Runs,
			mcp.WithDescription(
				"List recent salt highstate runs. Returns metadata without full output.",
			),
			mcp.WithNumber(
				constant.Limit,
				mcp.Description(
					"Maximum number of runs to return. Default: 20.",
				),
			),
		),
		s.runs,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Run,
			mcp.WithDescription(
				"Fetch a single salt highstate run by ID with full output.",
			),
			mcp.WithNumber(
				constant.Identifier,
				mcp.Required(),
				mcp.Description("Run ID."),
			),
		),
		s.getRun,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Minions,
			mcp.WithDescription(
				"List connected salt minions.",
			),
		),
		s.minions,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Keys,
			mcp.WithDescription(
				"List salt minion key status: accepted, pending, denied, rejected.",
			),
		),
		s.keys,
	)
}
