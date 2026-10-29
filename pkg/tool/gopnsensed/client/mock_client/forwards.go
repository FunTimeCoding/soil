package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) Forwards(
	_ context.Context,
	_ string,
) ([]client.Forward, error) {
	return c.ForwardList, c.Failure
}
