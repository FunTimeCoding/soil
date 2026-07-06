package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/network"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
)

func (c *Client) MustAssignPhysicalToInterface(
	p *physical_address.Address,
	i *network.Interface,
) *physical_address.Address {
	result, e := c.AssignPhysicalToInterface(p, i)
	errors.PanicOnError(e)

	return result
}
