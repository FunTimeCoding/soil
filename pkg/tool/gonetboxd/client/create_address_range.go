package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateAddressRange(
	start string,
	end string,
	status string,
	description string,
) (string, int) {
	body := client.CreateAddressRangeRequest{Start: start, End: end}

	if status != "" {
		body.Status = &status
	}

	if description != "" {
		body.Description = &description
	}

	result, e := c.client.CreateAddressRange(c.context, body)
	errors.PanicOnError(e)

	return web.ReadString(result), result.StatusCode
}
