package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostNotify(
	x context.Context,
	body client.PostNotifyJSONRequestBody,
) (*client.DeliveryResponse, error) {
	r, e := c.generated.PostNotifyWithResponse(x, body)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post notify", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
