package salt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response/local_return"
)

func (c *Client) MustLocal(
	glob string,
	function string,
	a []string,
) map[string]local_return.LocalReturn {
	result, e := c.Local(glob, function, a)
	errors.PanicOnError(e)

	return result
}
