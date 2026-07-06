package netbox

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteInterface(identifier int32) {
	errors.PanicOnError(c.DeleteInterface(identifier))
}
