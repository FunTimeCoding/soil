package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerOrganization() {
	s.server.AddTool(
		mcp.NewTool(
			constant.Whoami,
			mcp.WithDescription("Get the authenticated Sentry user"),
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
			mcp.WithDescription("List releases in the configured organization"),
			mcp.WithString("query", mcp.Description("Filter query")),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results"),
			),
		),
		mcp.NewTypedToolHandler(s.FindReleases),
	)
}
