package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostTurnEnd(
	x context.Context,
	body client.PostTurnEndJSONRequestBody,
) error {
	r, e := c.generated.PostTurnEndWithResponse(x, body)

	if e != nil {
		return e
	}

	if r.StatusCode() < http.StatusOK || r.StatusCode() >= http.StatusMultipleChoices {
		return refusal.New("post turn end", r.HTTPResponse, r.Body)
	}

	return nil
}
