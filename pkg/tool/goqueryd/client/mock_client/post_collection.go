package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) PostCollection(
	_ context.Context,
	b client.PostCollectionJSONRequestBody,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return c.Failure
	}

	c.CollectionsAdded = append(c.CollectionsAdded, b)

	return nil
}
