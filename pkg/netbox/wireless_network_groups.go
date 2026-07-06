package netbox

import "github.com/funtimecoding/soil/pkg/netbox/wireless_network_group"

func (c *Client) WirelessNetworkGroups() ([]*wireless_network_group.Group, error) {
	result, _, e := c.client.WirelessAPI.WirelessWirelessLanGroupsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return wireless_network_group.NewSlice(result.Results), nil
}
