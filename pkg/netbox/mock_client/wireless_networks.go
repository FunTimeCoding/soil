package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/wireless_network"

func (c *Client) WirelessNetworks() ([]*wireless_network.Network, error) {
	return nil, nil
}
