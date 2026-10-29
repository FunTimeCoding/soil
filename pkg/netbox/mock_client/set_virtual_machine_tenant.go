package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (c *Client) SetVirtualMachineTenant(
	_ string,
	_ string,
) (*virtual_machine.Machine, error) {
	return nil, nil
}
