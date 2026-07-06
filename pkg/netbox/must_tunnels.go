package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/tunnel"
)

func (c *Client) MustTunnels() []*tunnel.Tunnel {
	result, e := c.Tunnels()
	errors.PanicOnError(e)

	return result
}
