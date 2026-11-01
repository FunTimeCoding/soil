package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) SeedJob(
	identifier int64,
	name string,
	stage string,
	status string,
) {
	c.jobs = append(
		c.jobs,
		job.New(
			&gitlab.Job{
				ID:     identifier,
				Name:   name,
				Stage:  stage,
				Status: status,
			},
		),
	)
}
