package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostBackfill(
	x context.Context,
	params *client.PostBackfillParams,
) (*client.BackfillResponse, error) {
	r, e := c.generated.PostBackfillWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post backfill", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
