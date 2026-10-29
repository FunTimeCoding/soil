package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetConfigurationContext(
	x context.Context,
	name string,
) (client.ConfigurationContext, error) {
	r, e := c.client.GetConfigurationContextWithResponse(x, name)

	if e != nil {
		return client.ConfigurationContext{}, e
	}

	if r.JSON200 == nil {
		return client.ConfigurationContext{}, refusal.New(
			"get configuration context",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
