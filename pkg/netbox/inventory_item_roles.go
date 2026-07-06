package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item_role"
)

func (c *Client) InventoryItemRoles() ([]*inventory_item_role.Role, error) {
	result, _, e := c.client.DcimAPI.DcimInventoryItemRolesList(
		c.context,
	).Limit(constant.PageLimit).Execute()

	if e != nil {
		return nil, e
	}

	return inventory_item_role.NewSlice(result.Results), nil
}
