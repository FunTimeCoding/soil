package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/wireless_network"

func (c *Client) CreateWirelessNetwork(_ string) (*wireless_network.Network, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return nil, nil
}
