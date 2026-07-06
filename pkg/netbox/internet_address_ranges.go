package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) InternetAddressRanges() ([]netbox.IPRange, error) {
	result, _, e := c.client.IpamAPI.IpamIpRangesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return result.Results, nil
}
