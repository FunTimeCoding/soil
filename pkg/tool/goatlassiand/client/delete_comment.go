package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) DeleteComment(
	x context.Context,
	key string,
	identifier string,
) error {
	r, e := c.client.DeleteCommentWithResponse(x, key, identifier)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("delete comment", r.HTTPResponse, r.Body)
	}

	return nil
}
