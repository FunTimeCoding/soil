package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/device_type"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) DeviceTypes() ([]*device_type.Type, error) {
	if len(c.cache.DeviceTypes) != 0 {
		return c.cache.DeviceTypes, nil
	}

	result, _, e := c.client.DcimAPI.DcimDeviceTypesList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	c.cache.DeviceTypes = device_type.NewSlice(result.Results)

	return c.cache.DeviceTypes, nil
}
