package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/wireless_link"
)

func (c *Client) MustWirelessLinks() []*wireless_link.Link {
	result, e := c.WirelessLinks()
	errors.PanicOnError(e)

	return result
}
