package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"net"
)

func (c *Client) PhysicalExists(a net.HardwareAddr) (bool, error) {
	_, e := c.PhysicalAddress(a)

	if e != nil {
		if not_found.Is(e) {
			return false, nil
		}

		return false, e
	}

	return true, nil
}
