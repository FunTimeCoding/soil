package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) DeleteEntry(
	x context.Context,
	identifier int,
) error {
	r, e := c.client.DeleteEntryWithResponse(x, identifier)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("delete entry", r.HTTPResponse, r.Body)
	}

	return nil
}
