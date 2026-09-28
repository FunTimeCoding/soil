package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
)

func (c *Client) Item(identifier string) (*item.Item, error) {
	for _, i := range c.items {
		if i.Identifier == identifier {
			return i, nil
		}
	}

	return nil, fmt.Errorf("item %s not found", identifier)
}
