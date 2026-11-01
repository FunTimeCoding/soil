package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) pulseWith(
	sessionIdentifier string,
	body string,
	d client.Delivery,
) (bool, error) {
	r, e := c.PostSessionPulse(
		context.Background(),
		sessionIdentifier,
		client.PostSessionPulseJSONRequestBody{Body: body, Delivery: &d},
	)

	if e != nil {
		return false, e
	}

	return r.Immediate, nil
}
