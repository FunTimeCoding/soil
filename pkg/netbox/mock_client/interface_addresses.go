package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/internet_address"

func (c *Client) InterfaceAddresses(_ int32) ([]*internet_address.Address, error) {
	return nil, c.Failure
}
