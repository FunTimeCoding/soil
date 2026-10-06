package hub

import (
	"github.com/funtimecoding/soil/pkg/docker/hub/tag"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustTags(image string) []*tag.Tag {
	result, e := c.Tags(image)
	errors.PanicOnError(e)

	return result
}
