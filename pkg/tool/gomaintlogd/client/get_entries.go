package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetEntries(
	x context.Context,
	p *client.GetEntriesParams,
) ([]client.EntryResponse, error) {
	r, e := c.client.GetEntriesWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("entries", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
