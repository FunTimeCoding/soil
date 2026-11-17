package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) CreatePowerCable(
	_ context.Context,
	_ string,
	_ string,
	_ string,
	_ string,
) (client.Cable, error) {
	return client.Cable{}, c.Failure
}
