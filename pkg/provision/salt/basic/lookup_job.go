package basic

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) LookupJob(jobIdentifier string) (*response.Job, error) {
	var r response.JobDetail

	if e := c.Get(
		fmt.Sprintf("%s/%s", constant.SaltJobsPath, jobIdentifier),
		&r,
	); e != nil {
		return nil, e
	}

	if len(r.Details) == 0 {
		return nil, not_found.New("job", jobIdentifier)
	}

	job := r.Details[0]

	if job.Error != "" {
		return nil, unexpected.Format("job %s: %s", jobIdentifier, job.Error)
	}

	return &job, nil
}
