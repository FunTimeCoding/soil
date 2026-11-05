package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (c *Client) SetVirtualMachinePrimaryAddress(
	_ string,
	_ string,
) (*virtual_machine.Machine, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
