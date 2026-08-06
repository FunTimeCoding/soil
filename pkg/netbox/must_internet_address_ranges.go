package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address_range"
)

func (c *Client) MustInternetAddressRanges() []*internet_address_range.Range {
	result, e := c.InternetAddressRanges()
	errors.PanicOnError(e)

	return result
}
