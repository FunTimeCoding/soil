package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostSessionContext(
	x context.Context,
	identifier string,
	body client.PostSessionContextJSONRequestBody,
) error {
	r, e := c.generated.PostSessionContextWithResponse(x, identifier, body)

	if e != nil {
		return e
	}

	if r.StatusCode() < http.StatusOK || r.StatusCode() >= http.StatusMultipleChoices {
		return refusal.New("post session context", r.HTTPResponse, r.Body)
	}

	return nil
}
