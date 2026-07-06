package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network"
)

func (c *Client) MustVirtualNetworks() []*virtual_network.Network {
	result, e := c.VirtualNetworks()
	errors.PanicOnError(e)

	return result
}
