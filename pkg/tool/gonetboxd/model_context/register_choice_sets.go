package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerChoiceSets() {
	name := mcp.WithString(
		generativeConstant.ParameterName,
		mcp.Required(),
		mcp.Description("Name"),
	)
	choices := mcp.WithArray(
		constant.Choices,
		mcp.Required(),
		mcp.WithStringItems(),
		mcp.Description("Every choice; replaces the old list"),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListChoiceSets,
			mcp.WithDescription("List NetBox custom field choice sets"),
		),
		mcp.NewTypedToolHandler(s.ListChoiceSets),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateChoiceSet,
			mcp.WithDescription("Create a NetBox custom field choice set"),
			name,
			choices,
		),
		mcp.NewTypedToolHandler(s.CreateChoiceSet),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateChoiceSet,
			mcp.WithDescription("Replace the choices of a NetBox choice set"),
			name,
			choices,
		),
		mcp.NewTypedToolHandler(s.UpdateChoiceSet),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteChoiceSet,
			mcp.WithDescription("Delete a NetBox custom field choice set"),
			name,
		),
		mcp.NewTypedToolHandler(s.DeleteChoiceSet),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListCustomFields,
			mcp.WithDescription("List NetBox custom fields"),
		),
		mcp.NewTypedToolHandler(s.ListCustomFields),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateCustomField,
			mcp.WithDescription(
				"Create a multiple-selection NetBox custom field on devices and virtual machines, choosing from a choice set",
			),
			name,
			mcp.WithString(
				constant.ChoiceSet,
				mcp.Required(),
				mcp.Description("Choice set the field selects from"),
			),
			mcp.WithString(constant.Label, mcp.Description("Display label")),
		),
		mcp.NewTypedToolHandler(s.CreateCustomField),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteCustomField,
			mcp.WithDescription("Delete a NetBox custom field"),
			name,
		),
		mcp.NewTypedToolHandler(s.DeleteCustomField),
	)
}
