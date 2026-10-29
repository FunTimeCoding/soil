package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateAddressRange(
	x context.Context,
	start string,
	end string,
	status string,
	description string,
) (client.AddressRange, error) {
	body := client.CreateAddressRangeRequest{Start: start, End: end}

	if status != "" {
		body.Status = &status
	}

	if description != "" {
		body.Description = &description
	}

	r, e := c.client.CreateAddressRangeWithResponse(x, body)

	if e != nil {
		return client.AddressRange{}, e
	}

	if r.JSON201 == nil {
		return client.AddressRange{}, refusal.New(
			"create address range",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
