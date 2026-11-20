package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) DeleteSystem(
	x context.Context,
	name string,
) error {
	r, e := c.client.DeleteSystemWithResponse(x, name)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("delete system", r.HTTPResponse, r.Body)
	}

	return nil
}
