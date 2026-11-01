package model_context

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) registerInventory() {
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
}
