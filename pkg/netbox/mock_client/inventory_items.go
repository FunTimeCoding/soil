package mock_client

import "github.com/funtimecoding/soil/pkg/netbox/inventory_item"

func (c *Client) InventoryItems() ([]*inventory_item.Item, error) {
	return nil, c.Failure
}
