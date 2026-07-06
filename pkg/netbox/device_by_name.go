package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) DeviceByName(n string) (*device.Device, error) {
	result, e := c.DevicesByName(n)

	if e != nil {
		return nil, e
	}

	if len(result) > 1 {
		for _, r := range result {
			if r.Name == n {
				return r, nil
			}
		}

		return nil, fmt.Errorf(
			"no exact match for device %s among %d results: %w",
			n,
			len(result),
			constant.ErrorNotFound,
		)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf(
			"device not found: %s: %w",
			n,
			constant.ErrorNotFound,
		)
	}

	return result[0], nil
}
