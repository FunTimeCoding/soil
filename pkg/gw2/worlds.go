package gw2

import (
	"github.com/MrGunflame/gw2api"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Worlds() []*gw2api.World {
	result, e := c.client.Worlds()
	errors.PanicOnError(e)

	return result
}
