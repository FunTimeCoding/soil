package netbox

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeletePhysical(identifier int32) {
	errors.PanicOnError(c.DeletePhysical(identifier))
}
