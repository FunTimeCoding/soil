package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) GetTokens(
	_ context.Context,
	p *client.GetTokensParams,
) (*client.TokenSummary, error) {
	c.TokenCalls = append(c.TokenCalls, p)

	if c.Failure != nil {
		return nil, c.Failure
	}

	if c.Tokens != nil {
		return c.Tokens, nil
	}

	return &client.TokenSummary{}, nil
}
