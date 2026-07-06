package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network_group"
)

func (c *Client) MustWirelessNetworkGroups() []*wireless_network_group.Group {
	result, e := c.WirelessNetworkGroups()
	errors.PanicOnError(e)

	return result
}
