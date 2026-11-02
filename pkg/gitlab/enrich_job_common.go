package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
)

func (c *Client) enrichJobCommon(
	x context.Context,
	j *job.Job,
) error {
	if j.Fail() {
		trace, e := c.Trace(x, j.Project.Identifier, j.Identifier)

		if e != nil {
			return e
		}

		j.Trace = trace
	}

	j.Validate()

	return nil
}
