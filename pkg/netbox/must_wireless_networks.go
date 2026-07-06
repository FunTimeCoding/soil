package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network"
)

func (c *Client) MustWirelessNetworks() []*wireless_network.Network {
	result, e := c.WirelessNetworks()
	errors.PanicOnError(e)

	return result
}
