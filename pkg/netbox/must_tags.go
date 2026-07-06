package netbox

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/netbox/tag"
)

func (c *Client) MustTags() []*tag.Tag {
	result, e := c.Tags()
	errors.PanicOnError(e)

	return result
}
