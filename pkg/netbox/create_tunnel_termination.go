package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateTunnelTermination(
	t *tunnel.Tunnel,
	terminationType string,
	terminationIdentifier int64,
	role string,
) (*tunnel_termination.Termination, error) {
	q := netbox.NewWritableTunnelTerminationRequest(
		netbox.Int32AsPatchedWritableTunnelTerminationRequestTunnel(&t.Identifier),
		terminationType,
	)
	q.SetTerminationId(terminationIdentifier)
	q.SetRole(
		netbox.PatchedWritableTunnelTerminationRequestRole(role),
	)
	result, _, e := c.client.VpnAPI.VpnTunnelTerminationsCreate(
		c.context,
	).WritableTunnelTerminationRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	return tunnel_termination.New(result), nil
}
