package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
)

func (c *Client) Search(
	_ context.Context,
	_ *client.SearchParams,
) (*client.SearchResult, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.SearchResult), nil
}
