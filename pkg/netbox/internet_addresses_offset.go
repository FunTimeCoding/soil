package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) internetAddressesOffset(offset int32) ([]netbox.IPAddress, error) {
	result, _, e := c.client.IpamAPI.IpamIpAddressesList(
		c.context,
	).Limit(constant.PageLimit).Offset(offset).Execute()

	if e != nil {
		return nil, e
	}

	return result.Results, nil
}
