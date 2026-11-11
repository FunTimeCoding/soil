package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/smartctl/response"
)

func (c *Client) Drives(_ context.Context) ([]*response.Report, error) {
	c.Calls++

	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.drives, nil
}
