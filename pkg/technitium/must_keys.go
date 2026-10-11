package technitium

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/technitium/key"
)

func (c *Client) MustKeys() []*key.Key {
	result, e := c.Keys()
	errors.PanicOnError(e)

	return result
}
