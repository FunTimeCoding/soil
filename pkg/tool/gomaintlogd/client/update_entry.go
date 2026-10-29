package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) UpdateEntry(
	x context.Context,
	identifier int,
	b client.PostEntryRequest,
) (*client.EntryResponse, error) {
	r, e := c.client.UpdateEntryWithResponse(x, identifier, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("update entry", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
