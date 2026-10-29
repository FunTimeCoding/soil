package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) DeleteHost(
	x context.Context,
	identifier string,
	apply bool,
) error {
	r, e := c.client.DeleteHostWithResponse(
		x,
		identifier,
		&client.DeleteHostParams{Apply: &apply},
	)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("delete host", r.HTTPResponse, r.Body)
	}

	return nil
}
