package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/internet_address"

func (c *Client) CreateVirtualAddress(
	_ int32,
	_ string,
	_ string,
) (*internet_address.Address, error) {
	return nil, nil
}
