package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/custom_field"

func (c *Client) CustomFields() ([]*custom_field.Field, error) {
	return c.fields, nil
}
