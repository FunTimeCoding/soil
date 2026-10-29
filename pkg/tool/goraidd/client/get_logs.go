package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetLogs(
	x context.Context,
	p *client.GetLogsParams,
) (client.LogsResponse, error) {
	r, e := c.client.GetLogsWithResponse(x, p)

	if e != nil {
		return client.LogsResponse{}, e
	}

	if r.JSON200 == nil {
		return client.LogsResponse{}, refusal.New(
			"logs",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
