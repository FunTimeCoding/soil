package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateTunnel(
	name string,
	encapsulation string,
	group *tunnel_group.Group,
) (*tunnel.Tunnel, error) {
	q := netbox.NewWritableTunnelRequest(
		name,
		netbox.PatchedWritableTunnelRequestEncapsulation(encapsulation),
	)
	q.SetGroup(
		netbox.BriefTunnelGroupRequestAsPatchedWritableTunnelRequestGroup(
			netbox.NewBriefTunnelGroupRequest(
				group.Name,
				group.Raw.GetSlug(),
			),
		),
	)
	q.SetStatus(netbox.PATCHEDWRITABLETUNNELREQUESTSTATUS_ACTIVE)
	result, _, e := c.client.VpnAPI.VpnTunnelsCreate(
		c.context,
	).WritableTunnelRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return tunnel.New(result), nil
}
