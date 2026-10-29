package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) TransitionIssue(
	x context.Context,
	key string,
	transitionIdentifier string,
) error {
	r, e := c.client.TransitionIssueWithResponse(
		x,
		key,
		client.TransitionRequest{TransitionIdentifier: transitionIdentifier},
	)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("transition issue", r.HTTPResponse, r.Body)
	}

	return nil
}
