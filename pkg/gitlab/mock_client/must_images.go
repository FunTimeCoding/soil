package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gitlab/image"
)

func (c *Client) MustImages(
	x context.Context,
	project int64,
	repository int64,
) []*image.Image {
	result, e := c.Images(x, project, repository)
	errors.PanicOnError(e)

	return result
}
