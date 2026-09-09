package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/tag"

func (c *Client) Tags(_ int64) ([]*tag.Tag, error) {
	return nil, nil
}
