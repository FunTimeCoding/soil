package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_machine"

func (c *Client) SetVirtualMachineCores(
	_ string,
	_ float64,
) (*virtual_machine.Machine, error) {
	return nil, nil
}
