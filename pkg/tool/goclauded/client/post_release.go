package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostRelease(
	x context.Context,
	body client.PostReleaseJSONRequestBody,
) error {
	r, e := c.generated.PostReleaseWithResponse(x, body)

	if e != nil {
		return e
	}

	if r.StatusCode() < http.StatusOK || r.StatusCode() >= http.StatusMultipleChoices {
		return refusal.New("post release", r.HTTPResponse, r.Body)
	}

	return nil
}
