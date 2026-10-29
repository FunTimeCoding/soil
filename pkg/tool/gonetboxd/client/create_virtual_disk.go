package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) CreateVirtualDisk(
	x context.Context,
	machine string,
	name string,
	size int32,
) (client.VirtualDisk, error) {
	r, e := c.client.CreateVirtualDiskWithResponse(
		x,
		machine,
		client.CreateVirtualDiskRequest{Name: name, Size: size},
	)

	if e != nil {
		return client.VirtualDisk{}, e
	}

	if r.JSON201 == nil {
		return client.VirtualDisk{}, refusal.New(
			"create virtual disk",
			r.HTTPResponse,
			r.Body,
		)
	}

	return *r.JSON201, nil
}
