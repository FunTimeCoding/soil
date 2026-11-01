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
		s.listChoiceSets,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.CreateChoiceSet,
			mcp.WithDescription("Create a NetBox custom field choice set"),
			name,
			choices,
		),
		s.createChoiceSet,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.UpdateChoiceSet,
			mcp.WithDescription("Replace the choices of a NetBox choice set"),
			name,
			choices,
		),
		s.updateChoiceSet,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteChoiceSet,
			mcp.WithDescription("Delete a NetBox custom field choice set"),
			name,
		),
		s.deleteChoiceSet,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ListCustomFields,
			mcp.WithDescription("List NetBox custom fields"),
		),
		s.listCustomFields,
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
		s.createCustomField,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.DeleteCustomField,
			mcp.WithDescription("Delete a NetBox custom field"),
			name,
		),
		s.deleteCustomField,
	)
}
