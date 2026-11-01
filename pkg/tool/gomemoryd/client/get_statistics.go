package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetStatistics(x context.Context) (*client.Statistics, error) {
	r, e := c.client.GetStatisticsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("statistics", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
