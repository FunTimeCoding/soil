package netbox

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteInternet(identifier int32) {
	errors.PanicOnError(c.DeleteInternet(identifier))
}
