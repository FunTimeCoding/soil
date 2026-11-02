package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetLinkNotes(
	x context.Context,
	identifier int32,
	p *client.GetLinkNotesParams,
) (*client.NotePage, error) {
	r, e := c.client.GetLinkNotesWithResponse(x, identifier, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get link notes", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
