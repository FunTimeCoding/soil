package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) PhysicalAddresses() ([]*physical_address.Address, error) {
	if len(c.cache.PhysicalAddresses) > 0 {
		return c.cache.PhysicalAddresses, nil
	}

	var result []netbox.MACAddress
	var offset int32

	for {
		page, e := c.physicalAddressesOffset(offset)

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
		return nil, fmt.Errorf("no physical addresses found")
	}

	c.cache.PhysicalAddresses = physical_address.NewSlice(result)

	return c.cache.PhysicalAddresses, nil
}
