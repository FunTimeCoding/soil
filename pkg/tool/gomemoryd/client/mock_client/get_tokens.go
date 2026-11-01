package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) GetTokens(
	context.Context,
	*client.GetTokensParams,
) (*client.TokenSummary, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.TokenSummary{}, nil
}
