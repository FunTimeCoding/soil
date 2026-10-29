package mock_client

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) SeedSizedDeviceType(
	model string,
	height float64,
) int32 {
	v := netbox.NewDeviceTypeWithDefaults()
	v.SetId(c.nextIdentifier())
	v.SetModel(model)
	v.SetUHeight(height)
	c.types = append(c.types, device_type.New(v))

	return v.GetId()
}
