package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/response"

func (c *Client) GraphRunner(_ int64) (*response.Runner, error) {
	return nil, c.failure
}
