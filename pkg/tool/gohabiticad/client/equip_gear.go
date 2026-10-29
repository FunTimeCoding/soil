package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) EquipGear(
	x context.Context,
	key string,
) (*client.GearResult, error) {
	r, e := c.client.EquipGearWithResponse(x, key)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("equip gear", r.HTTPResponse, r.Body)
	}

	return r.JSON200, nil
}
