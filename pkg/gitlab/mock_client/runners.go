package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/runner"

func (c *Client) Runners(_ bool) ([]*runner.Runner, error) {
	return nil, c.failure
}
