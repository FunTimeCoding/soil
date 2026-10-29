package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/contact"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) Contacts() ([]*contact.Contact, error) {
	result, _, e := c.client.TenancyAPI.TenancyContactsList(c.context).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return contact.NewSlice(result.Results), nil
}
