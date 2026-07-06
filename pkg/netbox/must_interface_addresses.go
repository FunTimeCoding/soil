package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
)

func (c *Client) MustInterfaceAddresses(interfaceIdentifier int32) []*internet_address.Address {
	result, e := c.InterfaceAddresses(interfaceIdentifier)
	errors.PanicOnError(e)

	return result
}
