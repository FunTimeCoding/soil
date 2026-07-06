package salt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) MustMinions() []response.Minion {
	result, e := c.Minions()
	errors.PanicOnError(e)

	return result
}
