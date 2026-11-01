package model_context

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerIssues() {
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchIssues,
			mcp.WithDescription("Search issues in the configured organization"),
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
			mcp.WithString("cursor", mcp.Description("Pagination cursor")),
		),
		mcp.NewTypedToolHandler(s.SearchIssues),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.GetIssue,
			mcp.WithDescription("Get a Sentry issue by ID or short ID"),
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
			mcp.WithDescription("Get tag values for a Sentry issue"),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Issue ID or short ID"),
			),
			mcp.WithString(
				soil.TagKey,
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
}
