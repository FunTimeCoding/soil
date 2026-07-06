package netbox

import "github.com/funtimecoding/soil/pkg/netbox/custom_field_choice"

func (c *Client) CustomFieldChoices() ([]*custom_field_choice.Choice, error) {
	result, _, e := c.client.ExtrasAPI.ExtrasCustomFieldChoiceSetsList(
		c.context,
	).Execute()

	if e != nil {
		return nil, e
	}

	return custom_field_choice.NewSlice(result.Results), nil
}
