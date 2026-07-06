package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
)

func (c *Client) MustUpdatePhysicalAddress(
	a *physical_address.Address,
	i *network.Interface,
) *physical_address.Address {
	result, e := c.UpdatePhysicalAddress(a, i)
	errors.PanicOnError(e)

	return result
}
