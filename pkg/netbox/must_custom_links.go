package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/custom_link"
)

func (c *Client) MustCustomLinks() []*custom_link.Link {
	result, e := c.CustomLinks()
	errors.PanicOnError(e)

	return result
}
