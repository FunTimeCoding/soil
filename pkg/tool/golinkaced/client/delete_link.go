package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) DeleteLink(
	x context.Context,
	identifier int32,
) error {
	r, e := c.client.DeleteLinkWithResponse(x, identifier)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusOK {
		return refusal.New("delete link", r.HTTPResponse, r.Body)
	}

	return nil
}
