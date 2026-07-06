package netbox

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeviceTagNames(name string) []string {
	result, e := c.DeviceTagNames(name)
	errors.PanicOnError(e)

	return result
}
