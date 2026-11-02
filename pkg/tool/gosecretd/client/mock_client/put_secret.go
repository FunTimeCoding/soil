package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"maps"
	"slices"
)

func (c *Client) PutSecret(
	_ context.Context,
	p *client.PutSecretParams,
	b client.PutSecretJSONRequestBody,
) (*client.SecretKeys, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	merged := maps.Clone(c.Secrets[p.Path])

	if merged == nil {
		merged = map[string]string{}
	}

	maps.Copy(merged, b.Values)
	c.Secrets[p.Path] = merged

	return &client.SecretKeys{Keys: slices.Sorted(maps.Keys(merged))}, nil
}
