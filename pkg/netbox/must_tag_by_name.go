package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
)

func (c *Client) MustTagByName(n string) *tag.Tag {
	result, e := c.TagByName(n)
	errors.PanicOnError(e)

	return result
}
