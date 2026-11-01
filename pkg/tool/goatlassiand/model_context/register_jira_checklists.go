package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerJiraChecklists() {
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetChecklist,
			mcp.WithDescription(
				"Get the Smart Checklist items on a Jira issue. Returns indexed items with checked/unchecked status.",
			),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
		),
		s.getChecklist,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraAddChecklistItem,
			mcp.WithDescription(
				"Add an item to a Jira issue's Smart Checklist. Added as unchecked.",
			),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				constant.Text,
				mcp.Required(),
				mcp.Description("Checklist item text"),
			),
		),
		s.addChecklistItem,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraToggleChecklistItem,
			mcp.WithDescription(
				"Toggle a Smart Checklist item between checked and unchecked.",
			),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithNumber(
				constant.Index,
				mcp.Required(),
				mcp.Description("Item index (1-based)"),
			),
		),
		s.toggleChecklistItem,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraEditChecklistItem,
			mcp.WithDescription("Edit the text of a Smart Checklist item."),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithNumber(
				constant.Index,
				mcp.Required(),
				mcp.Description("Item index (1-based)"),
			),
			mcp.WithString(
				constant.Text,
				mcp.Required(),
				mcp.Description("New item text"),
			),
		),
		s.editChecklistItem,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraDeleteChecklistItem,
			mcp.WithDescription(
				"Delete an item from a Jira issue's Smart Checklist.",
			),
			mcp.WithString(
				generativeConstant.ParameterKey,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithNumber(
				constant.Index,
				mcp.Required(),
				mcp.Description("Item index (1-based)"),
			),
		),
		s.deleteChecklistItem,
	)
}
