package gitlab

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/image"
)

func (c *Client) MustImages(
	project int64,
	repository int64,
) []*image.Image {
	result, e := c.Images(project, repository)
	errors.PanicOnError(e)

	return result
}
