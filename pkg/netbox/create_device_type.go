package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/manufacturer"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateDeviceType(
	model string,
	m *manufacturer.Manufacturer,
) (*device_type.Type, error) {
	q := netbox.NewWritableDeviceTypeRequest(
		netbox.BriefManufacturerRequestAsBriefDeviceTypeRequestManufacturer(
			netbox.NewBriefManufacturerRequest(m.Name, m.Raw.Slug),
		),
		model,
		slug(model),
	)
	result, _, e := c.client.DcimAPI.DcimDeviceTypesCreate(
		c.context,
	).WritableDeviceTypeRequest(*q).Execute()

	if e != nil {
		return nil, e
	}

	c.cache.DeviceTypes = nil

	return device_type.New(result), nil
}
