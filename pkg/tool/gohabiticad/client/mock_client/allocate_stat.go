package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
)

func (c *Client) AllocateStat(
	_ context.Context,
	_ string,
) (*client.Stats, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.StatsValue, nil
}
