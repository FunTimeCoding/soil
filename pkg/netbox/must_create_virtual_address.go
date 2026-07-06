package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
)

func (c *Client) MustCreateVirtualAddress(
	interfaceIdentifier int32,
	address string,
) *internet_address.Address {
	result, e := c.CreateVirtualAddress(interfaceIdentifier, address)
	errors.PanicOnError(e)

	return result
}
