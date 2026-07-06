package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"net"
)

func (c *Client) MustCreatePhysicalInterface(
	a net.HardwareAddr,
	description string,
	i *network.Interface,
) *physical_address.Address {
	result, e := c.CreatePhysicalInterface(a, description, i)
	errors.PanicOnError(e)

	return result
}
