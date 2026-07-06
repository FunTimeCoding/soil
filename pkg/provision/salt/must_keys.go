package salt

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/salt/basic/response"
)

func (c *Client) MustKeys() *response.KeysReturn {
	result, e := c.Keys()
	errors.PanicOnError(e)

	return result
}
