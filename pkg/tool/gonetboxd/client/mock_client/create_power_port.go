package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreatePowerPort(
	_ context.Context,
	_ string,
	_ string,
) (client.PowerPort, error) {
	return client.PowerPort{}, c.Failure
}
