package mock_client

import "github.com/funtimecoding/soil/pkg/habitica/tag"

func (c *Client) Tags() ([]*tag.Tag, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.tags, nil
}
