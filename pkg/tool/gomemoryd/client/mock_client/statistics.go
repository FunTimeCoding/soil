package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) Statistics(context.Context) (*client.Statistics, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Stats, nil
}
