package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"strings"
)

func (c *Client) DevicesByMatch(s string) ([]*device.Device, error) {
	var result []*device.Device

	for _, d := range c.devices {
		if strings.Contains(d.Name, s) {
			result = append(result, d)
		}
	}

	return result, nil
}
