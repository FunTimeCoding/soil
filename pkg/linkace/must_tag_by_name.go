package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

func (c *Client) MustTagByName(name string) *tag.Tag {
	result, e := c.TagByName(name)
	errors.PanicOnError(e)

	return result
}
