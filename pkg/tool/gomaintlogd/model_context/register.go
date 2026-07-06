package model_context

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.AddEntry,
			mcp.WithDescription("Add a new maintenance log entry"),
			mcp.WithString(
				constant.Action,
				mcp.Required(),
				mcp.Description(
					"Brief description of what was done (e.g., 'restarted service', 'deployed v1.2.3')",
				),
			),
			mcp.WithString(
				constant.User,
				mcp.Required(),
				mcp.Description(
					"Who performed the action (username, email, or identifier)",
				),
			),
			mcp.WithString(
				constant.System,
				mcp.Description(
					"System identifier (node name, hostname, IP address) - optional",
				),
			),
			mcp.WithString(
				constant.Service,
				mcp.Description(
					"Service identifier (project name, deployment, application) - optional",
				),
			),
			mcp.WithString(
				constant.Description,
				mcp.Description(
					"Additional details or context about the action - optional",
				),
			),
			mcp.WithString(
				constant.Timestamp,
				mcp.Description(
					"When the event happened (RFC3339 format) - defaults to now if not provided",
				),
			),
		),
		s.add,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListEntries,
			mcp.WithDescription(
				"List maintenance log entries with optional filters",
			),
			mcp.WithString(
				constant.System,
				mcp.Description("Filter by system identifier"),
			),
			mcp.WithString(
				constant.Service,
				mcp.Description("Filter by service identifier"),
			),
			mcp.WithString(
				constant.User,
				mcp.Description("Filter by user"),
			),
			mcp.WithString(
				constant.Since,
				mcp.Description(
					"Show entries since this time (RFC3339 format)",
				),
			),
			mcp.WithString(
				constant.Until,
				mcp.Description(
					"Show entries until this time (RFC3339 format)",
				),
			),
			mcp.WithNumber(
				parameter.Limit,
				mcp.Description(
					"Maximum number of entries to return (default: all)",
				),
			),
		),
		s.list,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateEntry,
			mcp.WithDescription(
				"Update an existing maintenance log entry",
			),
			mcp.WithNumber(
				constant.Identifier,
				mcp.Required(),
				mcp.Description("ID of the entry to update"),
			),
			mcp.WithString(
				constant.Action,
				mcp.Description("Updated action description"),
			),
			mcp.WithString(
				constant.User,
				mcp.Description("Updated user"),
			),
			mcp.WithString(
				constant.System,
				mcp.Description("Updated system identifier"),
			),
			mcp.WithString(
				constant.Service,
				mcp.Description("Updated service identifier"),
			),
			mcp.WithString(
				constant.Description,
				mcp.Description("Updated description"),
			),
			mcp.WithString(
				constant.Timestamp,
				mcp.Description(
					"Updated timestamp (RFC3339 format)",
				),
			),
		),
		s.update,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteEntry,
			mcp.WithDescription(
				"Delete a maintenance log entry",
			),
			mcp.WithNumber(
				constant.Identifier,
				mcp.Required(),
				mcp.Description("ID of the entry to delete"),
			),
		),
		s.delete,
	)
}
