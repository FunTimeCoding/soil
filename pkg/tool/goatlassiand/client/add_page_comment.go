package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) AddPageComment(
	x context.Context,
	identifier string,
	body string,
) error {
	r, e := c.client.AddPageCommentWithResponse(
		x,
		identifier,
		client.CommentRequest{Body: body},
	)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("add page comment", r.HTTPResponse, r.Body)
	}

	return nil
}
