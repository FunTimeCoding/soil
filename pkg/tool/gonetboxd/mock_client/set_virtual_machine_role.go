package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (c *Client) SetVirtualMachineRole(
	name string,
	roleName string,
) (*virtual_machine.Machine, error) {
	return c.changeMachine("SetVirtualMachineRole", name, roleName)
}
