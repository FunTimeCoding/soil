package netbox

import "github.com/funtimecoding/soil/pkg/netbox/contact"

func (c *Client) Contacts() ([]*contact.Contact, error) {
	result, _, e := c.client.TenancyAPI.TenancyContactsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return contact.NewSlice(result.Results), nil
}
