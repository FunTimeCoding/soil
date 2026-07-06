package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/contact"
)

func (c *Client) MustContacts() []*contact.Contact {
	result, e := c.Contacts()
	errors.PanicOnError(e)

	return result
}
