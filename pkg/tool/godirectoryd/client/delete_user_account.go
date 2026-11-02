package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) DeleteUserAccount(
	x context.Context,
	account string,
) (*client.Receipt, error) {
	r, e := c.client.DeleteUserAccountWithResponse(x, account)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("delete user account", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
