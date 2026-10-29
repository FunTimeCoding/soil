package mock_client

import "github.com/funtimecoding/soil/pkg/kubernetes/types/native/job"

func (c *Client) Job(_ string, _ string) (*job.Job, error) {
	return nil, c.Failure
}
