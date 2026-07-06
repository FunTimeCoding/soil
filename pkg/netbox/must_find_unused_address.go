package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net"
)

func (c *Client) MustFindUnusedAddress(sub *net.IPNet) net.IP {
	result, e := c.FindUnusedAddress(sub)
	errors.PanicOnError(e)

	return result
}
