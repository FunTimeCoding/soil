package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerEvents() {
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
			mcp.WithDescription("Search events for a specific Sentry issue"),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Issue ID or short ID"),
			),
			mcp.WithString("query", mcp.Description("Search query")),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results"),
			),
			mcp.WithString("cursor", mcp.Description("Pagination cursor")),
		),
		mcp.NewTypedToolHandler(s.SearchIssueEvents),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SearchEvents,
			mcp.WithDescription(
				"Search events across the configured organization",
			),
			mcp.WithString("query", mcp.Description("Search query")),
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
		mcp.NewTypedToolHandler(s.SearchEvents),
	)
}
