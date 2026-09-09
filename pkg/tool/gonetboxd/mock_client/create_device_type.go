package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
)

func (c *Client) CreateDeviceType(
	_ string,
	_ *manufacturer.Manufacturer,
) (*device_type.Type, error) {
	return nil, nil
}
