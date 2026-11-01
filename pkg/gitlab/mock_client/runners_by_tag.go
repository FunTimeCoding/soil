package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/runner"

func (c *Client) RunnersByTag(_ string) ([]*runner.Runner, error) {
	return nil, c.failure
}
