package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (c *Client) PipelineJobs(
	_ int64,
	_ int64,
) ([]*job.Job, error) {
	return nil, nil
}
