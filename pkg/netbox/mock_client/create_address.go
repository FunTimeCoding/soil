package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/internet_address"

func (c *Client) CreateAddress(
	_ int32,
	_ string,
	_ string,
) (*internet_address.Address, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
