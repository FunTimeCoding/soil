package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) Blocklists(
	_ context.Context,
	_ string,
) ([]client.Blocklist, error) {
	return c.BlocklistList, c.Failure
}
