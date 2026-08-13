package basic

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) LookupJob(jobIdentifier string) (*response.Job, error) {
	b, e := c.Get(fmt.Sprintf("%s/%s", constant.SaltJobsPath, jobIdentifier))

	if e != nil {
		return nil, e
	}

	var r response.JobDetail

	if f := json.Unmarshal(b, &r); f != nil {
		return nil, f
	}

	if len(r.Details) == 0 {
		return nil, fmt.Errorf("job %s not found", jobIdentifier)
	}

	job := r.Details[0]

	if job.Error != "" {
		return nil, fmt.Errorf("job %s: %s", jobIdentifier, job.Error)
	}

	return &job, nil
}
