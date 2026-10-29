package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
)

func (c *Client) GetGear(context.Context) (*client.GearResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.GearValue, nil
}
