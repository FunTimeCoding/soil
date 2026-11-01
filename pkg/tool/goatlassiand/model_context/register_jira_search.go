package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerJiraSearch() {
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraSearch,
			mcp.WithDescription("Search Jira issues using JQL"),
			mcp.WithString(
				generativeConstant.ParameterQuery,
				mcp.Required(),
				mcp.Description("JQL query string"),
			),
			mcp.WithNumber(
				generativeConstant.ParameterLimit,
				mcp.Required(),
				mcp.Description("Maximum number of results (e.g. 25)"),
			),
		),
		s.searchIssues,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetIssue,
			mcp.WithDescription("Get a Jira issue by key"),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key (e.g. PROJ-123)"),
			),
			mcp.WithBoolean(
				constant.CustomFields,
				mcp.Description(
					"Include custom field values relevant to the issue's project and type. Requires an additional API call.",
				),
			),
			mcp.WithBoolean(
				constant.Comments,
				mcp.Description("Include comments on the issue."),
			),
		),
		s.getIssue,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraSearchProjects,
			mcp.WithDescription("Search Jira projects by key or name"),
			mcp.WithString(
				generativeConstant.ParameterQuery,
				mcp.Description(
					"Filter by key (exact, case-insensitive) or name (case-insensitive contains). Omit to list all.",
				),
			),
			mcp.WithNumber(
				generativeConstant.ParameterLimit,
				mcp.Required(),
				mcp.Description("Maximum number of results (e.g. 25)"),
			),
			mcp.WithBoolean(
				constant.IncludeDescriptions,
				mcp.Description(
					"Fetch full project details including description (slower, one API call per result). Defaults to false.",
				),
			),
		),
		s.searchProjects,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraSearchUsers,
			mcp.WithDescription(
				"Search Jira users by display name or email. Returns account IDs needed for assigning issues.",
			),
			mcp.WithString(
				generativeConstant.ParameterQuery,
				mcp.Required(),
				mcp.Description("Search term (display name or email)"),
			),
		),
		s.searchUsers,
	)
}
