package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreatePrefix(
	x context.Context,
	prefix string,
	site string,
	description string,
) (client.Prefix, error) {
	body := client.CreatePrefixRequest{Prefix: prefix}

	if site != "" {
		body.Site = &site
	}

	if description != "" {
		body.Description = &description
	}

	r, e := c.client.CreatePrefixWithResponse(x, body)

	if e != nil {
		return client.Prefix{}, e
	}

	if r.JSON201 == nil {
		return client.Prefix{}, refusal.New(
			"create prefix",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
