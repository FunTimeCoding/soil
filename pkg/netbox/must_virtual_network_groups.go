package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_network_group"
)

func (c *Client) MustVirtualNetworkGroups() []*virtual_network_group.Group {
	result, e := c.VirtualNetworkGroups()
	errors.PanicOnError(e)

	return result
}
