package jenkins

import (
	"github.com/bndr/gojenkins"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Jobs() []*gojenkins.Job {
	result, e := c.client.GetAllJobs(c.context)
	errors.PanicOnError(e)

	return result
}
