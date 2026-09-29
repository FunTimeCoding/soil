package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
)

func (c *Client) MustCreateTag(name string) *tag.Tag {
	result, e := c.CreateTag(name)
	errors.PanicOnError(e)

	return result
}
