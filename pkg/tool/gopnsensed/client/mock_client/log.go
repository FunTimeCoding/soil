package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) Log(
	_ context.Context,
	_ int,
) ([]client.LogEntry, error) {
	return c.LogList, c.Failure
}
