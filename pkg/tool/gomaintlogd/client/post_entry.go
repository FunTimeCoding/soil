package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostEntry(
	x context.Context,
	b client.PostEntryRequest,
) (*client.EntryResponse, error) {
	r, e := c.client.PostEntryWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("add entry", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
