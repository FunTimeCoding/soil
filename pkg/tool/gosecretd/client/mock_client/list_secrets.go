package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
)

func (c *Client) ListSecrets(
	_ context.Context,
	_ *client.ListSecretsParams,
) (*client.Listing, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.Listing), nil
}
