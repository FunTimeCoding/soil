package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostUserAccountPassword(
	x context.Context,
	account string,
	b client.PostUserAccountPasswordJSONRequestBody,
) (*client.Receipt, error) {
	r, e := c.client.PostUserAccountPasswordWithResponse(x, account, b)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(
			"post user account password",
			r.HTTPResponse,
			r.Body,
		)
	}

	return r.JSON200, nil
}
