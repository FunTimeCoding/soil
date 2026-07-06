package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"net"
)

func (c *Client) MustCreatePhysical(
	a net.HardwareAddr,
	description string,
) *physical_address.Address {
	result, e := c.CreatePhysical(a, description)
	errors.PanicOnError(e)

	return result
}
