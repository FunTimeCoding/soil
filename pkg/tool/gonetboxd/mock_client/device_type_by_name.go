package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
)

func (c *Client) DeviceTypeByName(n string) (*device_type.Type, error) {
	for _, t := range c.types {
		if t.Model == n {
			return t, nil
		}
	}

	return nil, not_found.New(constant.DeviceTypeKind, n)
}
