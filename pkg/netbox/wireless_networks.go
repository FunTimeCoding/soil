package netbox

import "github.com/funtimecoding/soil/pkg/netbox/wireless_network"

func (c *Client) WirelessNetworks() ([]*wireless_network.Network, error) {
	result, _, e := c.client.WirelessAPI.WirelessWirelessLansList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return wireless_network.NewSlice(result.Results), nil
}
