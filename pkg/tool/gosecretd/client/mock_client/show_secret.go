package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"maps"
	"slices"
)

func (c *Client) ShowSecret(
	_ context.Context,
	p *client.ShowSecretParams,
) (*client.SecretKeys, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	values, e := c.secret(p.Path)

	if e != nil {
		return nil, e
	}

	return &client.SecretKeys{Keys: slices.Sorted(maps.Keys(values))}, nil
}
