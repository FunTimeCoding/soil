package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
)

func (c *Client) DeleteCollection(
	_ context.Context,
	p *client.DeleteCollectionParams,
) (bool, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return false, c.Failure
	}

	deleted := false

	for k, d := range c.documents {
		if d.Collection == p.Name {
			delete(c.documents, k)
			deleted = true
		}
	}

	return deleted, nil
}
