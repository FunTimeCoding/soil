package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/internet_address"

func (c *Client) CreateAddress(
	_ int32,
	_ string,
	_ string,
) (*internet_address.Address, error) {
	return nil, nil
}
