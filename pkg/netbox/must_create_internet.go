package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) MustCreateInternet(
	objectType string,
	objectIdentifier int64,
	i net.IP,
	m net.IPMask,
) *netbox.IPAddress {
	result, e := c.CreateInternet(objectType, objectIdentifier, i, m)
	errors.PanicOnError(e)

	return result
}
