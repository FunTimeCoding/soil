package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
)

func (c *Client) EquipGear(
	_ context.Context,
	key string,
) (*client.GearResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	c.GearValue.Equipped["weapon"] = key

	return c.GearValue, nil
}
