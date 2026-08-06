package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address_range"
)

func (c *Client) InternetAddressRanges() ([]*internet_address_range.Range, error) {
	result, _, e := c.client.IpamAPI.IpamIpRangesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return internet_address_range.NewSlice(result.Results), nil
}
