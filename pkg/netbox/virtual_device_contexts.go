package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/virtual_device_context"
)

func (c *Client) VirtualDeviceContexts() ([]*virtual_device_context.Context, error) {
	result, _, e := c.client.DcimAPI.DcimVirtualDeviceContextsList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return virtual_device_context.NewSlice(result.Results), nil
}
