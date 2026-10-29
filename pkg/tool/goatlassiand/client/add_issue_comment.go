package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) AddIssueComment(
	x context.Context,
	key string,
	body string,
) error {
	r, e := c.client.AddIssueCommentWithResponse(
		x,
		key,
		client.CommentRequest{Body: body},
	)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("add issue comment", r.HTTPResponse, r.Body)
	}

	return nil
}
