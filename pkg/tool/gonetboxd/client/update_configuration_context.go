package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) UpdateConfigurationContext(
	x context.Context,
	name string,
	body client.UpdateConfigurationContextRequest,
) (client.ConfigurationContext, error) {
	r, e := c.client.UpdateConfigurationContextWithResponse(x, name, body)

	if e != nil {
		return client.ConfigurationContext{}, e
	}

	if r.JSON200 == nil {
		return client.ConfigurationContext{}, refusal.New(
			"update configuration context",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
