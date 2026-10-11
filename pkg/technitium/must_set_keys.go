package technitium

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/technitium/key"
)

func (c *Client) MustSetKeys(v []*key.Key) {
	errors.PanicOnError(c.SetKeys(v))
}
