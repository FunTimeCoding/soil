package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) Hosts(
	_ context.Context,
	_ string,
) ([]client.Host, error) {
	return c.HostList, c.Failure
}
