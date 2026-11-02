package mock_client

import (
	"github.com/funtimecoding/soil/pkg/docker/hub/tag"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) Tags(image string) ([]*tag.Tag, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result, okay := c.tags[image]

	if !okay {
		return nil, not_found.New("image", image)
	}

	return result, nil
}
