package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) DeleteClaim(
	x context.Context,
	item string,
	owner string,
) error {
	r, e := c.client.DeleteClaimWithResponse(
		x,
		&client.DeleteClaimParams{Item: item, Owner: owner},
	)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("delete claim", r.HTTPResponse, r.Body)
	}

	return nil
}
