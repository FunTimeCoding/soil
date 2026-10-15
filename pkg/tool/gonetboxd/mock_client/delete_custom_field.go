package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/custom_field"
	"slices"
)

func (c *Client) DeleteCustomField(name string) error {
	before := len(c.fields)
	c.fields = slices.DeleteFunc(
		c.fields,
		func(x *custom_field.Field) bool { return x.Name == name },
	)

	if len(c.fields) == before {
		return not_found.New(constant.CustomFieldKind, name)
	}

	c.record("DeleteCustomField", name)

	return nil
}
