package gw2

import (
	"github.com/MrGunflame/gw2api"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Account() gw2api.Account {
	result, e := c.client.Account()
	errors.PanicOnError(e)

	return result
}
