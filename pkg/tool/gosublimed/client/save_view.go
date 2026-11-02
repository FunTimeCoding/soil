package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) SaveView(
	x context.Context,
	identifier int,
	b client.SaveViewJSONRequestBody,
) error {
	r, e := c.client.SaveViewWithResponse(x, identifier, b)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("save view", r.HTTPResponse, r.Body)
	}

	return nil
}
