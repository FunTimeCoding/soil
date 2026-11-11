package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/internet_address"

func (c *Client) VirtualMachineAddresses(machine string) ([]*internet_address.Address, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.addresses[machine], nil
}
