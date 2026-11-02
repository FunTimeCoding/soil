package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) EditNote(
	x context.Context,
	identifier int32,
	b client.EditNoteJSONRequestBody,
) (*client.Note, error) {
	r, e := c.client.EditNoteWithResponse(x, identifier, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("edit note", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
