package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) DeleteNote(
	x context.Context,
	identifier int32,
) error {
	r, e := c.client.DeleteNoteWithResponse(x, identifier)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusOK {
		return refusal.New("delete note", r.HTTPResponse, r.Body)
	}

	return nil
}
