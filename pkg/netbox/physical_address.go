package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"net"
)

func (c *Client) PhysicalAddress(a net.HardwareAddr) (*physical_address.Address, error) {
	result, e := c.PhysicalAddressesByHardware(a)

	if e != nil {
		return nil, e
	}

	if len(result) > 1 {
		for _, r := range result {
			if r.Address.String() == a.String() {
				return r, nil
			}
		}

		return nil, fmt.Errorf(
			"no exact match for physical address %s among %d results: %w",
			a,
			len(result),
			constant.ErrorNotFound,
		)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf(
			"physical address not found: %s: %w",
			a,
			constant.ErrorNotFound,
		)
	}

	return result[0], nil
}
