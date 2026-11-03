package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	jellyfinConstant "github.com/funtimecoding/soil/pkg/jellyfin/constant"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

func (c *Client) GetItem(
	_ context.Context,
	identifier string,
) (*client.Item, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	for _, i := range c.ItemList {
		if i.Id == identifier {
			return i, nil
		}
	}

	return nil, not_found.New(jellyfinConstant.ItemKind, identifier)
}
