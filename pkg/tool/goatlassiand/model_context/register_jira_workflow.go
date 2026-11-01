package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerJiraWorkflow() {
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraUpdateIssue,
			mcp.WithDescription(
				"Update an existing Jira issue. Only provided fields are changed. Empty strings are ignored, not applied. Returns the updated issue with a before/after diff of changed fields.",
			),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key (e.g. INF-123)"),
			),
			mcp.WithString(
				constant.Summary,
				mcp.Description("New summary/title"),
			),
			mcp.WithString("description", mcp.Description("New description")),
			mcp.WithString(
				constant.Assignee,
				mcp.Description(
					"Assignee display name or email. Resolved to account ID automatically. Pass \"none\" to unassign.",
				),
			),
			mcp.WithString(
				constant.Reporter,
				mcp.Description(
					"Reporter display name or email. Resolved to account ID automatically.",
				),
			),
			mcp.WithString(
				constant.Labels,
				mcp.Description(
					"JSON array of label strings. Replaces all existing labels.",
				),
			),
			mcp.WithString(
				constant.AdditionalFields,
				mcp.Description(
					"JSON object mapping field names to values for custom or optional fields. Use jira_get_create_meta to discover available field names and allowed values.",
				),
			),
			mcp.WithBoolean(
				constant.NoDiff,
				mcp.Description(
					"Skip the before/after diff in the response. Useful for large description changes.",
				),
			),
		),
		s.updateIssue,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetTransitions,
			mcp.WithDescription("List available transitions for a Jira issue"),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
		),
		s.getTransitions,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraTransitionIssue,
			mcp.WithDescription("Transition a Jira issue to a new status"),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				constant.TransitionIdentifier,
				mcp.Required(),
				mcp.Description("Transition ID from jira_get_transitions"),
			),
		),
		s.transitionIssue,
	)
}
