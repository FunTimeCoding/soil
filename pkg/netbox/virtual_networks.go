package netbox

import "github.com/funtimecoding/soil/pkg/netbox/virtual_network"

func (c *Client) VirtualNetworks() ([]*virtual_network.Network, error) {
	result, _, e := c.client.IpamAPI.IpamVlansList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return virtual_network.NewSlice(result.Results), nil
}
