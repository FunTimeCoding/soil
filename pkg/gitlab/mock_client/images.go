package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/image"
)

func (c *Client) Images(
	_ context.Context,
	_ int64,
	_ int64,
) ([]*image.Image, error) {
	return nil, c.failure
}
