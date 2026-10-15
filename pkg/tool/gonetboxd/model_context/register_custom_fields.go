package model_context

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerCustomFields() {
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
	field := mcp.WithString(
		constant.Field,
		mcp.Required(),
		mcp.Description("Custom field name"),
	)
	values := mcp.WithArray(
		constant.Values,
		mcp.Required(),
		mcp.WithStringItems(),
		mcp.Description("Selected values; empty clears the field"),
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.AnsibleInventory,
			mcp.WithDescription(
				"Ansible YAML inventory built from NetBox: every device and virtual machine with a value in the field, one group per value, ansible_host from the primary address, and the keys of the ansible object in the host's configuration context as host variables, overriding ansible_host",
			),
			mcp.WithString(
				constant.Field,
				mcp.Description(
					"Custom field whose values become the groups. Default: ansible_playbook.",
				),
			),
		),
		s.ansibleInventory,
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
	s.server.AddTool(
		mcp.NewTool(
			constant.SetDeviceCustomField,
			mcp.WithDescription("Set one custom field on a NetBox device"),
			mcp.WithString(
				constant.Device,
				mcp.Required(),
				mcp.Description("Device name"),
			),
			field,
			values,
		),
		s.setDeviceCustomField,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.SetVirtualCustomField,
			mcp.WithDescription(
				"Set one custom field on a NetBox virtual machine",
			),
			mcp.WithString(
				constant.VirtualMachine,
				mcp.Required(),
				mcp.Description("Virtual machine name"),
			),
			field,
			values,
		),
		s.setVirtualCustomField,
	)
}
