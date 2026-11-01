package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostSessionEnd(
	x context.Context,
	body client.PostSessionEndJSONRequestBody,
) error {
	r, e := c.generated.PostSessionEndWithResponse(x, body)

	if e != nil {
		return e
	}

	if r.StatusCode() < http.StatusOK || r.StatusCode() >= http.StatusMultipleChoices {
		return refusal.New("post session end", r.HTTPResponse, r.Body)
	}

	return nil
}
