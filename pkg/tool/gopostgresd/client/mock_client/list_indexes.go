package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) ListIndexes(
	_ context.Context,
	_ *client.ListIndexesParams,
) (*client.QueryResult, error) {
	return c.answer()
}
