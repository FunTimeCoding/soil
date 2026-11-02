package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) CloseView(
	x context.Context,
	identifier int,
) error {
	r, e := c.client.CloseViewWithResponse(x, identifier)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("close view", r.HTTPResponse, r.Body)
	}

	return nil
}
