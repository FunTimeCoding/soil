package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
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

	return nil, not_found.New("device type", n)
}
