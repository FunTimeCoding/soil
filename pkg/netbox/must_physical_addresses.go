package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
)

func (c *Client) MustPhysicalAddresses() []*physical_address.Address {
	result, e := c.PhysicalAddresses()
	errors.PanicOnError(e)

	return result
}
