package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
)

func (c *Client) CustomFieldByName(n string) (*custom_field.Field, error) {
	result, e := c.CustomFields()

	if e != nil {
		return nil, e
	}

	for _, v := range result {
		if v.Name == n {
			return v, nil
		}
	}

	return nil, not_found.New(constant.CustomFieldKind, n)
}
