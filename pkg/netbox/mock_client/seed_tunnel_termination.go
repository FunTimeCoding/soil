package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedTunnelTermination(
	terminationType string,
	identifier int64,
) *tunnel_termination.Termination {
	v := netbox.NewTunnelTerminationWithDefaults()
	v.SetTerminationType(terminationType)
	v.SetTerminationId(identifier)
	result := tunnel_termination.New(v)
	c.tunnels = append(c.tunnels, result)

	return result
}
