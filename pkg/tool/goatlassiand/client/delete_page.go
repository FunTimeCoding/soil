package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) DeletePage(
	x context.Context,
	identifier string,
	draft bool,
) error {
	params := &client.DeletePageParams{}

	if draft {
		params.Draft = &draft
	}

	r, e := c.client.DeletePageWithResponse(x, identifier, params)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("delete page", r.HTTPResponse, r.Body)
	}

	return nil
}
