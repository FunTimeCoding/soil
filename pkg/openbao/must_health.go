package openbao

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/openbao/status"
)

func (c *Client) MustHealth() *status.Status {
	result, e := c.Health()
	errors.PanicOnError(e)

	return result
}
