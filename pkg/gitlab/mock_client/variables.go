package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/variable"

func (c *Client) Variables(_ int64) ([]*variable.Variable, error) {
	return nil, nil
}
