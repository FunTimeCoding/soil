package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostEditSession(
	x context.Context,
	body client.PostEditSessionJSONRequestBody,
) error {
	r, e := c.generated.PostEditSessionWithResponse(x, body)

	if e != nil {
		return e
	}

	if r.StatusCode() < http.StatusOK || r.StatusCode() >= http.StatusMultipleChoices {
		return refusal.New("post edit session", r.HTTPResponse, r.Body)
	}

	return nil
}
