package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) PrefixAddresses(prefix string) ([]*internet_address.Address, error) {
	var result []netbox.IPAddress
	var offset int32

	for {
		page, e := c.prefixAddressesOffset(prefix, offset)

		if e != nil {
			return nil, e
		}

		result = append(result, page...)

		if len(page) < int(constant.PageLimit) {
			break
		}

		offset += constant.PageLimit
	}

	return internet_address.NewSlice(result), nil
}
