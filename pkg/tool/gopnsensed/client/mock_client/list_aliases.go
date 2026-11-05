package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) ListAliases(
	_ context.Context,
	_ string,
) ([]client.Alias, error) {
	return c.AliasList, c.Failure
}
