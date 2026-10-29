package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostTag(
	x context.Context,
	b client.PostTagJSONRequestBody,
) error {
	r, e := c.client.PostTagWithResponse(x, b)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusOK {
		return refusal.New("set tag", r.HTTPResponse, r.Body)
	}

	return nil
}
