package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
)

func (c *Client) MustInternetAddresses() []*internet_address.Address {
	result, e := c.InternetAddresses()
	errors.PanicOnError(e)

	return result
}
