package salt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) MustHighstate(target string) map[string]response.LocalReturn {
	result, e := c.Highstate(target)
	errors.PanicOnError(e)

	return result
}
