package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_termination"
)

func (c *Client) CreateTunnelTermination(
	_ *tunnel.Tunnel,
	_ string,
	_ int64,
	_ string,
) (*tunnel_termination.Termination, error) {
	return nil, nil
}
