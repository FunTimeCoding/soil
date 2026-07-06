package salt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) MustLocal(
	glob string,
	function string,
	a []string,
) map[string]response.LocalReturn {
	result, e := c.Local(glob, function, a)
	errors.PanicOnError(e)

	return result
}
