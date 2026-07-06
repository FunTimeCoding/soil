package gw2

import (
	"github.com/MrGunflame/gw2api"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Guild(identifier string) gw2api.Guild {
	result, e := c.client.Guild(identifier, true)
	errors.PanicOnError(e)

	return result
}
