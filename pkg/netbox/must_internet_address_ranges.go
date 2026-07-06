package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) MustInternetAddressRanges() []netbox.IPRange {
	result, e := c.InternetAddressRanges()
	errors.PanicOnError(e)

	return result
}
