package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PutSystem(
	x context.Context,
	name string,
	b client.SystemSnapshot,
) error {
	r, e := c.client.PutSystemWithResponse(x, name, b)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("put system", r.HTTPResponse, r.Body)
	}

	return nil
}
