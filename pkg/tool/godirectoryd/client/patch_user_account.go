package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PatchUserAccount(
	x context.Context,
	account string,
	b client.PatchUserAccountJSONRequestBody,
) (*client.User, error) {
	r, e := c.client.PatchUserAccountWithResponse(x, account, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("patch user account", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
