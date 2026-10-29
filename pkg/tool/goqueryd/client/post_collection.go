package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostCollection(
	x context.Context,
	b client.PostCollectionJSONRequestBody,
) error {
	r, e := c.client.PostCollectionWithResponse(x, b)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusOK {
		return refusal.New("add collection", r.HTTPResponse, r.Body)
	}

	return nil
}
