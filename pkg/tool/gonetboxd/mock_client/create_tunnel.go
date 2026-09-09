package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
)

func (c *Client) CreateTunnel(
	_ string,
	_ string,
	_ *tunnel_group.Group,
) (*tunnel.Tunnel, error) {
	return nil, nil
}
