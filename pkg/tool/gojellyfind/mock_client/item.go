package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/jellyfin/constant"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
)

func (c *Client) Item(identifier string) (*item.Item, error) {
	for _, i := range c.items {
		if i.Identifier == identifier {
			return i, nil
		}
	}

	return nil, not_found.New(constant.ItemKind, identifier)
}
