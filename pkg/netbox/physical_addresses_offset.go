package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) physicalAddressesOffset(offset int32) ([]netbox.MACAddress, error) {
	result, _, e := c.client.DcimAPI.DcimMacAddressesList(
		c.context,
	).Limit(constant.PageLimit).Offset(offset).Execute()

	if e != nil {
		return nil, e
	}

	return result.Results, nil
}
