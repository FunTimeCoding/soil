package mock_client

import "github.com/funtimecoding/soil/pkg/github/run"

func (c *Client) LatestRuns(_ string, _ string) ([]*run.Run, error) {
	return nil, c.Failure
}
