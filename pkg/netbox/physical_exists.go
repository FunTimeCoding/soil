package netbox

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"net"
)

func (c *Client) PhysicalExists(a net.HardwareAddr) (bool, error) {
	_, e := c.PhysicalAddress(a)

	if e != nil {
		if errors.Is(e, constant.ErrorNotFound) {
			return false, nil
		}

		return false, e
	}

	return true, nil
}
