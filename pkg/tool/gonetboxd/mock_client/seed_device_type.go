package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedDeviceType(model string) {
	v := netbox.NewDeviceTypeWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetModel(model)
	c.types = append(c.types, device_type.New(v))
}
