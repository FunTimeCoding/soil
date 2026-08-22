package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (c *Client) DeviceByNames(n []string) (*device.Device, error) {
	var result *device.Device

	for _, name := range n {
		devices, e := c.DevicesByName(name)

		if e != nil {
			return nil, e
		}

		if len(devices) == 0 {
			continue
		}

		if len(devices) > 1 {
			var identifiers []string

			for _, d := range devices {
				identifiers = append(identifiers, fmt.Sprintf("%d", d.Raw.Id))
			}

			return nil, ambiguous.Format(
				"more than one device named %s: %s",
				name,
				join.Comma(identifiers),
			)
		}

		result = devices[0]

		break
	}

	if result == nil {
		return nil, not_found.Format(
			"no device found matching names: %s",
			join.Comma(n),
		)
	}

	return result, nil
}
