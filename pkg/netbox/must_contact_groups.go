package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/contact_group"
)

func (c *Client) MustContactGroups() []*contact_group.Group {
	result, e := c.ContactGroups()
	errors.PanicOnError(e)

	return result
}
