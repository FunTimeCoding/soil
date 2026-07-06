package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/tag"

func (c *Client) Tags() ([]*tag.Tag, error) {
	return c.tags, nil
}
