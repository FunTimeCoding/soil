package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
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

		return nil, not_found.Format(
			"no exact match for device %s among %d results",
			n,
			len(result),
		)
	}

	if len(result) == 0 {
		return nil, not_found.New("device", n)
	}

	return result[0], nil
}
