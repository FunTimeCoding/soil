package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
)

func (c *Client) GetStatus(_ context.Context) (*client.Status, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.Status), nil
}
