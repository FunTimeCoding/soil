package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/internet_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) InternetAddresses() ([]*internet_address.Address, error) {
	if len(c.cache.InternetAddresses) > 0 {
		return c.cache.InternetAddresses, nil
	}

	var result []netbox.IPAddress
	var offset int32

	for {
		page, e := c.internetAddressesOffset(offset)

		if e != nil {
			return nil, e
		}

		result = append(result, page...)

		if len(page) < int(constant.PageLimit) {
			break
		}

		offset += constant.PageLimit
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no internet addresses found")
	}

	c.cache.InternetAddresses = internet_address.NewSlice(result)

	return c.cache.InternetAddresses, nil
}
