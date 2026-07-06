package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
)

func (c *Client) DeviceTypeByName(n string) (*device_type.Type, error) {
	result, e := c.DeviceTypes()

	if e != nil {
		return nil, e
	}

	for _, t := range result {
		if t.Model == n {
			return t, nil
		}
	}

	return nil, fmt.Errorf("device type not found: %s", n)
}
