package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PutSecret(
	x context.Context,
	p *client.PutSecretParams,
	b client.PutSecretJSONRequestBody,
) (*client.SecretKeys, error) {
	r, e := c.client.PutSecretWithResponse(x, p, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("put secret", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
