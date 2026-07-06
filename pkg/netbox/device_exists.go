package netbox

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (c *Client) DeviceExists(name string) (bool, error) {
	_, e := c.DeviceByName(name)

	if e != nil {
		if errors.Is(e, constant.ErrorNotFound) {
			return false, nil
		}

		return false, e
	}

	return true, nil
}
