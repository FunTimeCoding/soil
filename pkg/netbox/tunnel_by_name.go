package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
)

func (c *Client) TunnelByName(n string) (*tunnel.Tunnel, error) {
	result, e := c.Tunnels()

	if e != nil {
		return nil, e
	}

	for _, t := range result {
		if t.Name == n {
			return t, nil
		}
	}

	return nil, not_found.New("tunnel", n)
}
