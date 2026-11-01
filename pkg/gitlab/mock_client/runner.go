package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/runner"

func (c *Client) Runner(_ int64) (*runner.Runner, error) {
	return nil, c.failure
}
