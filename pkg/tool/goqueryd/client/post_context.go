package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostContext(
	x context.Context,
	b client.PostContextJSONRequestBody,
) error {
	r, e := c.client.PostContextWithResponse(x, b)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusOK {
		return refusal.New("add context", r.HTTPResponse, r.Body)
	}

	return nil
}
