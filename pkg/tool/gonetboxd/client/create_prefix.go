package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreatePrefix(
	prefix string,
	site string,
	description string,
) *response.Response {
	body := client.CreatePrefixRequest{Prefix: prefix}

	if site != "" {
		body.Site = &site
	}

	if description != "" {
		body.Description = &description
	}

	result, e := c.client.CreatePrefix(c.context, body)
	errors.PanicOnError(e)

	return response.New(web.ReadString(result), result.StatusCode)
}
