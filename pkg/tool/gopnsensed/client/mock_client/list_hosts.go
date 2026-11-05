package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) ListHosts(
	_ context.Context,
	_ string,
) ([]client.Host, error) {
	return c.HostList, c.Failure
}
