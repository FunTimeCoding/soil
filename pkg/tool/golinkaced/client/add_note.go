package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) AddNote(
	x context.Context,
	b client.AddNoteJSONRequestBody,
) (*client.Note, error) {
	r, e := c.client.AddNoteWithResponse(x, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("add note", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
