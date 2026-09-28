package mock_client

import "github.com/funtimecoding/soil/pkg/jellyfin/item"

func (c *Client) AddItem(i *item.Item) {
	c.items = append(c.items, i)
}
