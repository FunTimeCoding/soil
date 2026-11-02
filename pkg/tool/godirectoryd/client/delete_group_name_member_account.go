package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DeleteGroupNameMemberAccount(
	x context.Context,
	name string,
	account string,
) (*client.Group, error) {
	r, e := c.client.DeleteGroupNameMemberAccountWithResponse(x, name, account)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"delete group name member account",
			r.HTTPResponse,
			r.Body,
		)
	}

	return r.JSON200, nil
}
