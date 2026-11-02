package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
)

func (c *Client) SearchItems(
	_ context.Context,
	_ *client.SearchItemsParams,
) (*client.SearchResult, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.SearchResult), nil
}
