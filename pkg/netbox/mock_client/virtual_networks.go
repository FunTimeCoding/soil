package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/virtual_network"

func (c *Client) VirtualNetworks() ([]*virtual_network.Network, error) {
	return nil, c.Failure
}
