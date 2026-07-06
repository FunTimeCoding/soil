package netbox

import "github.com/funtimecoding/soil/pkg/netbox/custom_field"

func (c *Client) CustomFields() ([]*custom_field.Field, error) {
	result, _, e := c.client.ExtrasAPI.ExtrasCustomFieldsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return custom_field.NewSlice(result.Results), nil
}
