package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/custom_field"

func (c *Client) CustomFieldByName(_ string) (*custom_field.Field, error) {
	return nil, c.Failure
}
