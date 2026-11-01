package mock_client

import "github.com/funtimecoding/soil/pkg/github/run"

func (c *Client) FailedRuns(_ bool) ([]*run.Run, error) {
	return nil, c.Failure
}
