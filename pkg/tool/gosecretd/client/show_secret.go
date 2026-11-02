package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ShowSecret(
	x context.Context,
	p *client.ShowSecretParams,
) (*client.SecretKeys, error) {
	r, e := c.client.ShowSecretWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("show secret", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
