package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetUserAccount(
	x context.Context,
	account string,
) (*client.User, error) {
	r, e := c.client.GetUserAccountWithResponse(x, account)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get user account", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
