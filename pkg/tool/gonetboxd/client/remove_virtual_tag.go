package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) RemoveVirtualTag(
	x context.Context,
	name string,
	tag string,
) (client.VirtualMachine, error) {
	r, e := c.client.RemoveVirtualTagWithResponse(x, name, tag)

	if e != nil {
		return client.VirtualMachine{}, e
	}

	if r.JSON200 == nil {
		return client.VirtualMachine{}, refusal.New(
			"remove virtual tag",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON200, nil
}
