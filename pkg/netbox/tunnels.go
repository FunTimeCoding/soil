package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
)

func (c *Client) Tunnels() ([]*tunnel.Tunnel, error) {
	result, _, e := c.client.VpnAPI.VpnTunnelsList(c.context).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return tunnel.NewSlice(result.Results), nil
}
