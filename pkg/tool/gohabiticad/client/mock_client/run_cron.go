package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
)

func (c *Client) RunCron(context.Context) (*client.CronResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.CronValue, nil
}
