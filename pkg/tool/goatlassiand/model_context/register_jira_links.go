package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerJiraLinks() {
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetLinkTypes,
			mcp.WithDescription(
				"List available issue link types with their inward and outward labels.",
			),
		),
		s.getLinkTypes,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraLinkIssues,
			mcp.WithDescription(
				"Link two Jira issues. Use jira_get_link_types to discover available link type names.",
			),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Source issue key"),
			),
			mcp.WithString(
				constant.TargetKey,
				mcp.Required(),
				mcp.Description("Target issue key"),
			),
			mcp.WithString(
				constant.LinkType,
				mcp.Description(
					"Link type name (default: Relates). Use jira_get_link_types to see available names.",
				),
			),
		),
		s.linkIssues,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraDeleteLink,
			mcp.WithDescription(
				"Delete a link between two Jira issues. Use jira_get_issue to find link IDs.",
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Link ID"),
			),
		),
		s.deleteLink,
	)
}
