package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListSecrets(
	x context.Context,
	p *client.ListSecretsParams,
) (*client.Listing, error) {
	r, e := c.client.ListSecretsWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("list secrets", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
