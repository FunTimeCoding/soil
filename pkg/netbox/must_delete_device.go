package netbox

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteDevice(identifier int32) {
	errors.PanicOnError(c.DeleteDevice(identifier))
}
