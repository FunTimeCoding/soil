package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListPowerOutlets(
	_ context.Context,
	_ string,
) ([]*client.PowerOutlet, error) {
	return c.PowerOutletList, c.Failure
}
