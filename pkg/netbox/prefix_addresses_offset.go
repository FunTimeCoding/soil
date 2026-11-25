package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) prefixAddressesOffset(
	prefix string,
	offset int32,
) ([]netbox.IPAddress, error) {
	result, _, e := c.client.IpamAPI.IpamIpAddressesList(c.context).Parent(
		[]string{prefix},
	).Limit(
		constant.PageLimit,
	).Offset(
		offset,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return result.Results, nil
}
