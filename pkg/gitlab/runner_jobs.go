package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) RunnerJobs(
	runner int64,
	stopAfter int,
) ([]*job.Job, error) {
	var result []*gitlab.Job
	number := int64(1)

	for {
		page, _, e := c.client.Runners.ListRunnerJobs(
			runner,
			&gitlab.ListRunnerJobsOptions{
				ListOptions: gitlab.ListOptions{
					PerPage: constant.PerPage100,
					Page:    number,
				},
				OrderBy: new(constant.Identifier),
				Sort:    new(constant.Descending),
			},
		)

		if e != nil {
			return nil, e
		}

		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
			break
		}

		if stopAfter > 0 && len(result) >= stopAfter {
			result = job.Deduplicate(result)

			if len(result) >= stopAfter {
				break
			}
		}

		number++
	}

	return c.enrichJobs(job.NewSlice(result)), nil
}
