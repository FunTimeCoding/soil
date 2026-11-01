package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetTokens(
	x context.Context,
	p *client.GetTokensParams,
) (*client.TokenSummary, error) {
	r, e := c.client.GetTokensWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("tokens", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
