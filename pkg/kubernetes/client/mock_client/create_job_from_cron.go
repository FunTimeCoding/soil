package mock_client

import "k8s.io/api/batch/v1"

func (c *Client) CreateJobFromCron(
	_ string,
	_ string,
	_ string,
) (*v1.Job, error) {
	return nil, c.Failure
}
