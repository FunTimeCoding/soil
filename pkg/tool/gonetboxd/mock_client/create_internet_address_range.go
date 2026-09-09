package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/internet_address_range"

func (c *Client) CreateInternetAddressRange(
	_ string,
	_ string,
	_ string,
	_ string,
) (*internet_address_range.Range, error) {
	return nil, nil
}
