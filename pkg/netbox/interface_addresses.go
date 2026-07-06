package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
)

func (c *Client) InterfaceAddresses(interfaceIdentifier int32) ([]*internet_address.Address, error) {
	result, _, e := c.client.IpamAPI.IpamIpAddressesList(
		c.context,
	).Limit(constant.PageLimit).InterfaceId(
		[]int32{interfaceIdentifier},
	).Execute()

	if e != nil {
		return nil, e
	}

	return internet_address.NewSlice(result.Results), nil
}
