package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DeleteGroupName(
	x context.Context,
	name string,
) (*client.Receipt, error) {
	r, e := c.client.DeleteGroupNameWithResponse(x, name)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("delete group name", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
