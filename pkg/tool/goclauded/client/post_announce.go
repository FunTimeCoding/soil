package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostAnnounce(
	x context.Context,
	body client.PostAnnounceJSONRequestBody,
) error {
	r, e := c.generated.PostAnnounceWithResponse(x, body)

	if e != nil {
		return e
	}

	if r.StatusCode() < http.StatusOK || r.StatusCode() >= http.StatusMultipleChoices {
		return refusal.New("post announce", r.HTTPResponse, r.Body)
	}

	return nil
}
