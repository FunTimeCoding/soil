package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/image"

func (c *Client) Images(
	_ int64,
	_ int64,
) ([]*image.Image, error) {
	return nil, nil
}
