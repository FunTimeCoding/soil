package model_context

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Whoami,
			mcp.WithDescription(
				"Get the authenticated Sentry user",
			),
		),
		mcp.NewTypedToolHandler(s.Whoami),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.FindOrganizations,
			mcp.WithDescription("List all Sentry organizations"),
		),
		mcp.NewTypedToolHandler(s.FindOrganizations),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.FindProjects,
			mcp.WithDescription(
				"List all projects in the configured organization",
			),
		),
		mcp.NewTypedToolHandler(s.FindProjects),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.FindReleases,
			mcp.WithDescription(
				"List releases in the configured organization",
			),
			mcp.WithString(
				"query",
				mcp.Description("Filter query"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results"),
			),
		),
		mcp.NewTypedToolHandler(s.FindReleases),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchIssues,
			mcp.WithDescription(
				"Search issues in the configured organization",
			),
			mcp.WithString(
				"query",
				mcp.Description("Search query (e.g. is:unresolved)"),
			),
			mcp.WithString(
				"project",
				mcp.Description("Project slug to filter by"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results"),
			),
			mcp.WithString(
				"cursor",
				mcp.Description("Pagination cursor"),
			),
		),
		mcp.NewTypedToolHandler(s.SearchIssues),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetIssue,
			mcp.WithDescription(
				"Get a Sentry issue by ID or short ID",
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Issue ID or short ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetIssue),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetIssueTagValues,
			mcp.WithDescription(
				"Get tag values for a Sentry issue",
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Issue ID or short ID"),
			),
			mcp.WithString(
				library.TagKey,
				mcp.Required(),
				mcp.Description("Tag key"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results"),
			),
		),
		mcp.NewTypedToolHandler(s.GetIssueTagValues),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetIssueEvent,
			mcp.WithDescription(
				"Get the latest event for a Sentry issue, including request body, exception, and stack trace",
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Issue ID or short ID"),
			),
		),
		mcp.NewTypedToolHandler(s.GetIssueEvent),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetEvent,
			mcp.WithDescription(
				"Get a single event by its event ID, including exception and stack trace",
			),
			mcp.WithString(
				"project",
				mcp.Required(),
				mcp.Description("Project slug"),
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Event ID (32-character hex)"),
			),
		),
		mcp.NewTypedToolHandler(s.GetEvent),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchIssueEvents,
			mcp.WithDescription(
				"Search events for a specific Sentry issue",
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Issue ID or short ID"),
			),
			mcp.WithString(
				"query",
				mcp.Description("Search query"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results"),
			),
			mcp.WithString(
				"cursor",
				mcp.Description("Pagination cursor"),
			),
		),
		mcp.NewTypedToolHandler(s.SearchIssueEvents),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchEvents,
			mcp.WithDescription(
				"Search events across the configured organization",
			),
			mcp.WithString(
				"query",
				mcp.Description("Search query"),
			),
			mcp.WithString(
				"project",
				mcp.Description("Project slug to filter by"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results"),
			),
			mcp.WithString(
				"cursor",
				mcp.Description("Pagination cursor"),
			),
		),
		mcp.NewTypedToolHandler(s.SearchEvents),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateIssue,
			mcp.WithDescription(
				"Update the status or assignee of a Sentry issue",
			),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Issue ID or short ID"),
			),
			mcp.WithString(
				"status",
				mcp.Description(
					"New status (e.g. resolved, ignored, unresolved)",
				),
			),
			mcp.WithString(
				"assignedTo",
				mcp.Description("Username or team to assign to"),
			),
		),
		mcp.NewTypedToolHandler(s.UpdateIssue),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteIssue,
			mcp.WithDescription("Delete a Sentry issue"),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Issue ID or short ID"),
			),
		),
		mcp.NewTypedToolHandler(s.DeleteIssue),
	)
}
