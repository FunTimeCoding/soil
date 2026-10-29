package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) UpdateComment(
	x context.Context,
	key string,
	identifier string,
	body string,
) error {
	r, e := c.client.UpdateCommentWithResponse(
		x,
		key,
		identifier,
		client.CommentRequest{Body: body},
	)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("update comment", r.HTTPResponse, r.Body)
	}

	return nil
}
