package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateTunnelGroup(name string) (*tunnel_group.Group, error) {
	q := netbox.NewTunnelGroupRequest(name, slug(name))
	result, _, e := c.client.VpnAPI.VpnTunnelGroupsCreate(
		c.context,
	).TunnelGroupRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return tunnel_group.New(result), nil
}
