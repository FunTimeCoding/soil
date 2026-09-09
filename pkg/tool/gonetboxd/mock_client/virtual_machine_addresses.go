package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/internet_address"

func (c *Client) VirtualMachineAddresses(_ string) ([]*internet_address.Address, error) {
	return nil, nil
}
