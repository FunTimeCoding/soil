package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/physical_address"
	"net"
)

func (c *Client) EnsurePhysical(
	h net.HardwareAddr,
) (*physical_address.Address, error) {
	result, e := c.PhysicalAddressesByHardware(h)

	if e != nil {
		return nil, e
	}

	for _, r := range result {
		if r.Address.String() == h.String() {
			return r, nil
		}
	}

	return c.CreatePhysical(h, "")
}
