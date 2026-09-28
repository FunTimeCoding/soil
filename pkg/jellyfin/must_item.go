package jellyfin

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/jellyfin/item"
)

func (c *Client) MustItem(identifier string) *item.Item {
	result, e := c.Item(identifier)
	errors.PanicOnError(e)

	return result
}
