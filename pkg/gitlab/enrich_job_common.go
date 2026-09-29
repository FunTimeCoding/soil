package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/job"

func (c *Client) enrichJobCommon(j *job.Job) error {
	if j.Fail() {
		trace, e := c.Trace(j.Project.Identifier, j.Identifier)

		if e != nil {
			return e
		}

		j.Trace = trace
	}

	j.Validate()

	return nil
}
