package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/usage_result"
)

func (c *Client) Read(
	_ context.Context,
	_ string,
) (*usage_result.Result, error) {
	c.Reads++

	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Result, nil
}
