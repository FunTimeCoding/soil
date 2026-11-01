package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerJiraCreation() {
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetCreateMeta,
			mcp.WithDescription(
				"Get required and available fields for creating a Jira issue. Returns field names, IDs, whether each field is required, its schema type, and allowed values where applicable. Call this before jira_create_issue to discover what the target project and issue type expect.",
			),
			mcp.WithString(
				constant.Project,
				mcp.Required(),
				mcp.Description("Project key (e.g. OPS)"),
			),
			mcp.WithString(
				constant.IssueType,
				mcp.Required(),
				mcp.Description("Issue type name (e.g. Task, Bug, Story)"),
			),
			mcp.WithString(
				constant.ExpandFields,
				mcp.Description(
					"Comma-separated field names or keys to show full allowed values for. By default, fields with more than 5 allowed values are trimmed to the 5 most recent.",
				),
			),
		),
		s.getCreateMeta,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraCreateIssue,
			mcp.WithDescription(
				"Create a new Jira issue. Use jira_get_create_meta first to discover required fields and allowed values for the target project and issue type.",
			),
			mcp.WithString(
				constant.Project,
				mcp.Required(),
				mcp.Description("Project key (e.g. OPS)"),
			),
			mcp.WithString(
				constant.IssueType,
				mcp.Required(),
				mcp.Description("Issue type name (e.g. Task, Bug, Story)"),
			),
			mcp.WithString(
				constant.Summary,
				mcp.Required(),
				mcp.Description("Issue summary/title"),
			),
			mcp.WithString("description", mcp.Description("Issue description")),
			mcp.WithString(
				constant.Assignee,
				mcp.Description(
					"Assignee display name or email. Resolved to account ID automatically. Fails if no match or multiple matches.",
				),
			),
			mcp.WithString(
				constant.Labels,
				mcp.Description(
					"JSON array of label strings (e.g. [\"ZKM\",\"Infrastructure\"])",
				),
			),
			mcp.WithString(
				constant.AdditionalFields,
				mcp.Description(
					"JSON object mapping field names to values for custom or optional fields. Use jira_get_create_meta to discover available field names and allowed values.",
				),
			),
		),
		s.createIssue,
	)
}
