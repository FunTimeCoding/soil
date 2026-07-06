package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item_role"
)

func (c *Client) MustInventoryItemRoles() []*inventory_item_role.Role {
	result, e := c.InventoryItemRoles()
	errors.PanicOnError(e)

	return result
}
