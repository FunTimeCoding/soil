package netbox

import "github.com/funtimecoding/soil/pkg/netbox/virtual_network_group"

func (c *Client) VirtualNetworkGroups() ([]*virtual_network_group.Group, error) {
	result, _, e := c.client.IpamAPI.IpamVlanGroupsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return virtual_network_group.NewSlice(result.Results), nil
}
