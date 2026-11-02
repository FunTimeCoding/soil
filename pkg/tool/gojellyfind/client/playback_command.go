package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PlaybackCommand(
	x context.Context,
	identifier string,
	b client.PlaybackCommandJSONRequestBody,
) error {
	r, e := c.client.PlaybackCommandWithResponse(x, identifier, b)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("playback command", r.HTTPResponse, r.Body)
	}

	return nil
}
