package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"github.com/funtimecoding/soil/pkg/netbox/inventory_item"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) InventoryItems() ([]*inventory_item.Item, error) {
	result, _, e := c.client.DcimAPI.DcimInventoryItemsList(c.context).Limit(
		constant.PageLimit,
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return inventory_item.NewSlice(result.Results), nil
}
