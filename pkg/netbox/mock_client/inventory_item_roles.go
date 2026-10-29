package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/inventory_item_role"

func (c *Client) InventoryItemRoles() ([]*inventory_item_role.Role, error) {
	return nil, c.Failure
}
