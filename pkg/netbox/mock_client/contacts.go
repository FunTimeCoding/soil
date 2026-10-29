package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/contact"

func (c *Client) Contacts() ([]*contact.Contact, error) {
	return nil, c.Failure
}
