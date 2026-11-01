package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (c *Client) Jobs() ([]*job.Job, error) {
	return nil, c.failure
}
