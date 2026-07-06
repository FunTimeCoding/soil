package netbox

import (
	"fmt"
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

	return nil, fmt.Errorf("tunnel not found: %s", n)
}
