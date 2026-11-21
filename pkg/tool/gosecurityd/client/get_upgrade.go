package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetUpgrade(
	x context.Context,
	release string,
	source string,
) (*client.UpgradeGroup, error) {
	r, e := c.client.GetUpgradeWithResponse(x, release, source)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New(constant.UpgradeKind, r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
