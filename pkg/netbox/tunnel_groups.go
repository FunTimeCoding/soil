package netbox

import "github.com/funtimecoding/soil/pkg/netbox/tunnel_group"

func (c *Client) TunnelGroups() ([]*tunnel_group.Group, error) {
	result, _, e := c.client.VpnAPI.VpnTunnelGroupsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return tunnel_group.NewSlice(result.Results), nil
}
