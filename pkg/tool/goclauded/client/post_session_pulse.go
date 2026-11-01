package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostSessionPulse(
	x context.Context,
	identifier string,
	body client.PostSessionPulseJSONRequestBody,
) (*client.DeliveryResponse, error) {
	r, e := c.generated.PostSessionPulseWithResponse(x, identifier, body)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("post session pulse", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
