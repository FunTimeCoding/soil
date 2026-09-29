package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

func (c *Client) MustTags() []*tag.Tag {
	result, e := c.Tags()
	errors.PanicOnError(e)

	return result
}
