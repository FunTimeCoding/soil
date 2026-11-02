package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) RevealSecret(
	x context.Context,
	p *client.RevealSecretParams,
) (*client.SecretValue, error) {
	r, e := c.client.RevealSecretWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("reveal secret", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
