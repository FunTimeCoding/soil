package salt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response/local_return"
)

func (c *Client) MustHighstate(target string) map[string]local_return.LocalReturn {
	result, e := c.Highstate(target)
	errors.PanicOnError(e)

	return result
}
