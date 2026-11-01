package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerJiraComments() {
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraAddComment,
			mcp.WithDescription("Add a comment to a Jira issue"),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				generativeConstant.ParameterBody,
				mcp.Required(),
				mcp.Description("Comment text"),
			),
		),
		s.addIssueComment,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraUpdateComment,
			mcp.WithDescription(
				"Update an existing comment on a Jira issue. Use jira_get_issue with include_comments to find comment IDs.",
			),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Comment ID"),
			),
			mcp.WithString(
				generativeConstant.ParameterBody,
				mcp.Required(),
				mcp.Description("New comment body"),
			),
		),
		s.updateComment,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraDeleteComment,
			mcp.WithDescription(
				"Delete a comment from a Jira issue. Use jira_get_issue with include_comments to find comment IDs.",
			),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				generativeConstant.ParameterIdentifier,
				mcp.Required(),
				mcp.Description("Comment ID"),
			),
		),
		s.deleteComment,
	)
}
