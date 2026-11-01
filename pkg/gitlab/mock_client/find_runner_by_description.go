package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/runner"

func (c *Client) FindRunnerByDescription(_ string) (*runner.Runner, bool, error) {
	return nil, false, c.failure
}
