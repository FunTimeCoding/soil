package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) ListUpgrades(
	x context.Context,
	p *client.ListUpgradesParams,
) (*client.UpgradeList, error) {
	r, e := c.client.ListUpgradesWithResponse(x, p)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("upgrades", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
