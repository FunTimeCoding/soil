package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Playbooks,
			mcp.WithDescription(
				"List configured ansible playbook paths.",
			),
		),
		s.playbooks,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Sync,
			mcp.WithDescription(
				"Pull the latest changes from the repository. Returns a filtered log of commits that touched ansible playbooks.",
			),
		),
		s.syncRepository,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Trigger,
			mcp.WithDescription(
				"Trigger an ansible playbook run. Set update to pull latest code first. Set synchronous to wait for the result instead of returning immediately.",
			),
			mcp.WithString(
				constant.Playbook,
				mcp.Description(
					"Specific playbook path to run. Omit to run all.",
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
				"List recent ansible playbook runs. Returns metadata without full output.",
			),
			mcp.WithNumber(
				constant.Limit,
				mcp.Description("Maximum number of runs to return. Default: 20."),
			),
			mcp.WithString(
				constant.Status,
				mcp.Description(
					"Filter by status: pending, running, success, error.",
				),
			),
		),
		s.runs,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.Run,
			mcp.WithDescription(
				"Fetch a single ansible playbook run by ID with full output.",
			),
			mcp.WithNumber(
				constant.Identifier,
				mcp.Required(),
				mcp.Description("Run ID."),
			),
		),
		s.getRun,
	)
}
