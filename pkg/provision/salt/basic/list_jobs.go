package basic

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/soil/pkg/provision/salt/constant"
)

func (c *Client) ListJobs() ([]response.Job, error) {
	b, e := c.Get(constant.JobsPath)

	if e != nil {
		return nil, e
	}

	var r response.JobList

	if f := json.Unmarshal(b, &r); f != nil {
		return nil, f
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
