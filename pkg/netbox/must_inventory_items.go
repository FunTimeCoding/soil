package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item"
)

func (c *Client) MustInventoryItems() []*inventory_item.Item {
	result, e := c.InventoryItems()
	errors.PanicOnError(e)

	return result
}
