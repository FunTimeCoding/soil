package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) DeleteSession(
	x context.Context,
	callsign string,
) error {
	r, e := c.generated.DeleteSessionWithResponse(x, callsign)

	if e != nil {
		return e
	}

	if r.StatusCode() < http.StatusOK || r.StatusCode() >= http.StatusMultipleChoices {
		return refusal.New("delete session", r.HTTPResponse, r.Body)
	}

	return nil
}
