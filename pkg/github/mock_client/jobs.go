package mock_client

import "github.com/funtimecoding/soil/pkg/github/job"

func (c *Client) Jobs(_ string, _ string, _ int64) ([]*job.Job, error) {
	return nil, c.Failure
}
