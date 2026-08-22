package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
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

		return nil, not_found.Format(
			"no exact match for physical address %s among %d results",
			a,
			len(result),
		)
	}

	if len(result) == 0 {
		return nil, not_found.New("physical address", a)
	}

	return result[0], nil
}
