package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"net"
)

func (c *Client) MustHardwarePhysicalAddresses(a net.HardwareAddr) []*physical_address.Address {
	result, e := c.HardwarePhysicalAddresses(a)
	errors.PanicOnError(e)

	return result
}
