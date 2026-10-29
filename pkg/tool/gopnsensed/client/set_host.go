package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) SetHost(
	x context.Context,
	identifier string,
	h client.HostRequest,
	apply bool,
) error {
	r, e := c.client.SetHostWithResponse(
		x,
		identifier,
		&client.SetHostParams{Apply: &apply},
		h,
	)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("set host", r.HTTPResponse, r.Body)
	}

	return nil
}
