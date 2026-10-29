package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetReports(x context.Context) ([]client.ReportResponse, error) {
	r, e := c.client.GetReportsWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("reports", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
