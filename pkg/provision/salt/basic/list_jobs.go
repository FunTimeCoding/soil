package basic

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) ListJobs() ([]response.Job, error) {
	var r response.JobList

	if e := c.Get(constant.SaltJobsPath, &r); e != nil {
		return nil, e
	}

	if len(r.Return) == 0 {
		return nil, nil
	}

	var result []response.Job

	for jid, job := range r.Return[0] {
		job.JID = jid
		result = append(result, job)
	}

	return result, nil
}
