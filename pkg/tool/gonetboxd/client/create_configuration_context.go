package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateConfigurationContext(
	x context.Context,
	body client.CreateConfigurationContextRequest,
) (client.ConfigurationContext, error) {
	r, e := c.client.CreateConfigurationContextWithResponse(x, body)

	if e != nil {
		return client.ConfigurationContext{}, e
	}

	if r.JSON201 == nil {
		return client.ConfigurationContext{}, refusal.New(
			"create configuration context",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
