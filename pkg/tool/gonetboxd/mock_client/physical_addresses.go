package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/physical_address"

func (c *Client) PhysicalAddresses() ([]*physical_address.Address, error) {
	return nil, nil
}
