package offline_memory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) GetTokens(
	context.Context,
	*client.GetTokensParams,
) (*client.TokenSummary, error) {
	return nil, nil
}
