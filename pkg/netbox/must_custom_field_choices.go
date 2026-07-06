package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field_choice"
)

func (c *Client) MustCustomFieldChoices() []*custom_field_choice.Choice {
	result, e := c.CustomFieldChoices()
	errors.PanicOnError(e)

	return result
}
