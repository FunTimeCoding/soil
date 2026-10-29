package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) RunCron(x context.Context) (*client.CronResult, error) {
	r, e := c.client.RunCronWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("run cron", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
