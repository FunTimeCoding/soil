package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostGroupNameMember(
	x context.Context,
	name string,
	b client.PostGroupNameMemberJSONRequestBody,
) (*client.Group, error) {
	r, e := c.client.PostGroupNameMemberWithResponse(x, name, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"post group name member",
			r.HTTPResponse,
			r.Body,
		)
	}

	return r.JSON200, nil
}
