package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) SetVolume(
	x context.Context,
	identifier string,
	b client.SetVolumeJSONRequestBody,
) error {
	r, e := c.client.SetVolumeWithResponse(x, identifier, b)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("set volume", r.HTTPResponse, r.Body)
	}

	return nil
}
