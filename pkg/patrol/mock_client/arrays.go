package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/mdstat/array"
)

func (c *Client) Arrays(_ context.Context) ([]*array.Array, error) {
	c.Calls++

	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.arrays, nil
}
