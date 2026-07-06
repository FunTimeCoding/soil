package salt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) MustJob(identifier string) *response.Job {
	result, e := c.Job(identifier)
	errors.PanicOnError(e)

	return result
}
