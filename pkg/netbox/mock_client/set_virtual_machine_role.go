package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (c *Client) SetVirtualMachineRole(
	name string,
	roleName string,
) (*virtual_machine.Machine, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.changeMachine("SetVirtualMachineRole", name, roleName)
}
