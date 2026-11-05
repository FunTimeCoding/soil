package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
)

func (c *Client) FirewallStates(
	_ context.Context,
	_ string,
) ([]client.State, error) {
	return c.StateList, c.Failure
}
