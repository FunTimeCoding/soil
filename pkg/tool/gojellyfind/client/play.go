package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/constant"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) Play(
	x context.Context,
	identifier string,
	b client.PlayJSONRequestBody,
) error {
	r, e := c.client.PlayWithResponse(x, identifier, b)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New(constant.Play, r.HTTPResponse, r.Body)
	}

	return nil
}
