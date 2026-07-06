package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel_group"
)

func (c *Client) TunnelGroupByName(n string) (*tunnel_group.Group, error) {
	result, e := c.TunnelGroups()

	if e != nil {
		return nil, e
	}

	for _, g := range result {
		if g.Name == n {
			return g, nil
		}
	}

	return nil, fmt.Errorf("tunnel group not found: %s", n)
}
