package salt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) MustJobs() []response.Job {
	result, e := c.Jobs()
	errors.PanicOnError(e)

	return result
}
